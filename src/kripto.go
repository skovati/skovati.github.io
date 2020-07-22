package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
    "net/http"
    "io/ioutil"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "kripto.db")
	db.Exec("CREATE TABLE IF NOT EXISTS Portfolio (Currency TEXT, Amount FLOAT)")
	loop(db)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func addCrypto(db *sql.DB, currency string, amount float64) {
	tx, _ := db.Begin()
	statement, _ := tx.Prepare("INSERT INTO Portfolio (currency, amount) values (?,?)")
	_, err := statement.Exec(currency, amount)
	checkErr(err)
	tx.Commit()
}

func getPortfolio(db *sql.DB) {
    var sum float64
	rows, _ := db.Query("SELECT * FROM Portfolio ORDER BY Amount DESC")
	for rows.Next() {
		var tempCur string
		var tempAmt float64
		rows.Scan(&tempCur, &tempAmt)
        tempPrice := getPrice(tempCur)
        sum += tempPrice * tempAmt
		fmt.Printf("%s: %.3f, $%.2f\n", tempCur, tempAmt, tempPrice*tempAmt)
	}
    fmt.Println("")
    fmt.Printf("Total: $%.2f\n", sum)
    fmt.Println("---------------")
}

func removeCrypto(db *sql.DB, currency string) {
	tx, _ := db.Begin()
	statement, _ := tx.Prepare("DELETE FROM Portfolio WHERE Currency=?")
	_, err := statement.Exec(currency)
	checkErr(err)
	tx.Commit()
}

func updateCrypto(db *sql.DB, currency string, amount float64) {
	tx, _ := db.Begin()
	statement, _ := tx.Prepare("UPDATE Portfolio set Amount=? where Currency=?")
	_, err := statement.Exec(amount, currency)
	checkErr(err)
	tx.Commit()
}

func loop(db *sql.DB) {
	var input string = "1"
	fmt.Println("Welcome to Kripto!\n")
	options()
	fmt.Scanln(&input)
	for input != "0" {
		fmt.Println("")
		switch input {
		case "1":
			fmt.Println("Here is your current portfolio:", "\n---------------")
			getPortfolio(db)
			fmt.Println("")
		case "2":
			tempInput1, tempInput2 := dualInput("Enter the crypto you want to add: ", "Enter the amount you have: ")
			addCrypto(db, tempInput1, tempInput2)
		case "3":
			tempInput1, tempInput2 := dualInput("Enter the crypto you want to update: ", "Enter the new amount: ")
			updateCrypto(db, tempInput1, tempInput2)
		case "4":
			fmt.Print("Enter the crypto you want to remove: ")
			var tempInput1 string
			fmt.Scanln(&tempInput1)
			removeCrypto(db, tempInput1)
		}
		options()
		fmt.Scanln(&input)
	}
	fmt.Println("Come Again!")
}

func options() {
	fmt.Print("What would you like to do?\n",
		"0: Exit\n",
		"1: View Portfolio\n",
		"2: Add Crypto\n",
		"3: Update Crypto\n",
		"4: Remove Crypto\n\n",
		"Enter Choice: ")
}

func dualInput(prompt1, prompt2 string) (string, float64) {
	var input1 string
	var input2 string
	fmt.Print(prompt1)
	fmt.Scanln(&input1)
	fmt.Print("\n", prompt2)
	fmt.Scanln(&input2)
	floatInput, _ := strconv.ParseFloat(input2, 64)
	return input1, floatInput
}

func getPrice(currency string) float64 {
    currency = strings.ToLower(currency)
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + currency + "&vs_currencies=usd"
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]map[string]float64
	json.Unmarshal([]byte(body), &result)
	return result[currency]["usd"]
}
