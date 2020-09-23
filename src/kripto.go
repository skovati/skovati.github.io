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

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"

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
        fmt.Println(
            colorBlue, "\nKripto Portfolio:", colorReset, "\n\n",
            colorRed, "Coin", colorReset, "|",
            colorYellow, "Amount", colorReset, "|",
            colorGreen, "Value", colorReset, "|",
            colorPurple, "Market Price", colorReset, "|",
            colorCyan, "24h Ago", colorReset, "|",
            colorRed, "7d Ago", colorReset)
        fmt.Println("---------------------------------------------------------------------------")
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
    sum24 := 0.0
    sum7D := 0.0
    for s, v := range(portfolio) {
        price, price24, price7D, percent24, percent7D := getPrice(s)
        fmt.Printf("%s%s%s | %s%.3f%s | %s$%.2f%s | %s$%.2f%s | ", colorRed, strings.Title(s), colorReset,
            colorYellow, v, colorReset,
            colorGreen, v*price, colorReset,
            colorPurple, price, colorReset)
        if(percent24 < 0) {
            fmt.Printf("%s$%.2f%s(%s%.2f%%%s) | ", colorCyan, price24, colorReset, colorRed, percent24, colorReset)
        } else {
            fmt.Printf("%s$%.2f%s(%s%.2f%%%s) | ", colorCyan, price24, colorReset, colorGreen, percent24, colorReset)
        }
        if(percent7D < 0) {
            fmt.Printf("%s$%.2f%s(%s%.2f%%%s)\n", colorRed, price7D, colorReset, colorRed, percent7D, colorReset)
        } else {
            fmt.Printf("%s$%.2f%s(%s%.2f%%%s)\n", colorRed, price24, colorReset, colorGreen, percent7D, colorReset)
        }
        sum += price*v
        sum24 += price24*v
        sum7D += price7D*v
    }
    fmt.Println("---------------------------------------------------------------------------")
    fmt.Printf("\n%sTotal: $%.2f%s\n", colorGreen, sum, colorReset)

    totalPercent24 := ((sum/sum24)-1)*100
    totalPercent7D := ((sum/sum7D)-1)*100
    if(totalPercent24 < 0) {
        fmt.Printf("Total 24h ago: $%.2f(%s%.2f%%%s)\n", sum24, colorRed, totalPercent24, colorReset)
    } else {
        fmt.Printf("Total 24h ago: $%.2f(%s%.2f%%%s)\n", sum24, colorGreen, totalPercent24, colorReset)
    }
    if(totalPercent24 < 0) {
        fmt.Printf("Total 7d ago: $%.2f(%s%.2f%%%s)\n", sum7D, colorRed, totalPercent7D, colorReset)
    } else {
        fmt.Printf("Total 7d ago: $%.2f(%s%.2f%%%s)\n", sum7D, colorGreen, totalPercent7D, colorReset)
    }
}

func isErr(err error) bool {
	if err != nil {
        return true
	}
    return false
}

func getPrice(currency string) (usdPrice, usdPrice24, usdPrice7D, percentChange24, percentChange7D float64) {
    currency = strings.ToLower(currency)
    url := "https://api.coingecko.com/api/v3/coins/" + currency + "?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false"
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

	var result map[string]map[string]map[string]float64
	json.Unmarshal([]byte(body), &result)

    usdPrice = result["market_data"]["current_price"]["usd"]
    percentChange24 = result["market_data"]["price_change_percentage_24h_in_currency"]["usd"]
    percentChange7D = result["market_data"]["price_change_percentage_7d_in_currency"]["usd"]

    usdPrice24 = usdPrice / (1+(percentChange24/100))
    usdPrice7D = usdPrice / (1+(percentChange7D/100))

	return
}
