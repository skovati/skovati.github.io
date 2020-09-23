package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
    "net/http"
    "io/ioutil"
    "os"
)

/*
options
---------
init
add
remove
update
list
help
*/

var fileName string = findFileName()

func main() {
    if len(os.Args) < 2 {
        usage()
        return
    }
    switch os.Args[1] {
    case "help":
        usage()
    case "init":
        fmt.Println("Creating new portfolio...")
        newPortfolio()
        fmt.Println("New portfolio created, add currency using the add command.")
    case "add", "update":
        if len(os.Args) < 4 {
            incorrectUsage()
            return
        }
        amt, err := strconv.ParseFloat(os.Args[3], 64)
        cur := strings.ToLower(os.Args[2])
        if isErr(err) {
            fmt.Println("Please enter the amount of", os.Args[2], "you have")
        }
        fmt.Println("Adding", amt, "of", cur, "to your portfolio...")
        addCrypto(cur, amt)
        fmt.Println("Succesfully added", cur, "to your portfolio, use the list command to view.")
    case "remove":
        if len(os.Args) < 3 {
            incorrectUsage()
            return
        }
        var response string
        fmt.Println("Removing", os.Args[2], "from your portfolio... Are you sure? [Y/n]")
        fmt.Scan(&response)
        if response != "Y" {
            fmt.Println(os.Args[2], "has not been removed.")
            return
        }
        removeCrypto(strings.ToLower(os.Args[2]))
        fmt.Println(os.Args[2], "has been removed.")
    case "list":
        fmt.Println("\nKripto Portfolio:\nCoin, Amount, Value, Market Price")
        fmt.Println("-------------------------")
        listPortfolio()
        fmt.Println("")
    default:
        usage()
    }
}

func findFileName() string {
    fileName, err := os.UserHomeDir()
    if isErr(err) {
        fmt.Println("Cannot find an adequate location for portfolio...")
    }
    return fileName + "/.portfolio.json"
}

func usage() {
    fmt.Print(
        "\n",
        "Welcome to Kripto!\n\n",
        "commands:\n",
        "init - create a new portfolio\n",
        "add - Use 'kripto add <currency> <amount>' to add a new crypto\n",
        "remove - Use 'kripto remove <currency>' to remove\n",
        "update - Use 'kripto update <currency>' <new amount> to update holdings\n",
        "list - lists current porfolio and holdings\n",
        "help - show this message\n\n",
    )
}

func incorrectUsage() {
    fmt.Println("\nIncorrect usage, use the help command.")
    usage()
}

func newPortfolio() {
    _, err := os.Create(fileName)
    if isErr(err) {
        fmt.Println("Cannot create new portfolio.")
        return
    }
    return
}

func openFile() map[string]float64 {
    portfolio := make(map[string]float64)
    portfolioJson, err := ioutil.ReadFile(fileName)
    if isErr(err) {
        fmt.Println("File could not be read, try 'kripto init'")
        return nil
    }
    json.Unmarshal(portfolioJson, &portfolio)
    return portfolio
}

func savePortfolio(portfolio map[string]float64) {
    portfolioJson, err := json.MarshalIndent(portfolio, "", "   ")
    if isErr(err) {
        fmt.Println("Error encoding as json.")
    }
    err = ioutil.WriteFile(fileName, portfolioJson, 0770)
    if isErr(err) {
        fmt.Println("Error writing to file.")
    }
    return
}

func addCrypto(currency string, amount float64) {
    portfolio := openFile()
    portfolio[currency] = amount
    savePortfolio(portfolio)
}

func removeCrypto(currency string) {
    portfolio := openFile()
    delete(portfolio, currency)
    savePortfolio(portfolio)
}

func listPortfolio() {
    portfolio := openFile()
    if (len(portfolio)) == 0 {
        fmt.Println("Nothing to see here, add crypto using the add command.")
        return
    }
    sum := 0.0
    for s, v := range(portfolio) {
        price := getPrice(s)
        fmt.Printf("%s: %.3f, $%.2f, $%.2f\n", strings.Title(s), v, v*price, price)
        sum += price*v
    }
    fmt.Println("-------------------------")
    fmt.Printf("Total: $%.2f", sum)
    fmt.Println("")
}

func isErr(err error) bool {
	if err != nil {
        return true
	}
    return false
}

func getPrice(currency string) float64 {
    currency = strings.ToLower(currency)
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + currency + "&vs_currencies=usd"
	req, err := http.NewRequest("GET", url, nil)
	if isErr(err) {
        fmt.Println("err")
    }
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if isErr(err) {
        fmt.Println("err")
    }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]map[string]float64
	json.Unmarshal([]byte(body), &result)
	return result[currency]["usd"]
}
