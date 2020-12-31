package portfolio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
    "github.com/skovati/kripto/file"
    "github.com/skovati/kripto/coin"
)

// standard location for user data in Linux
const LINUX_PREFIX = "/.local/share/kripto"

// location of portfolio
var PortfolioDir string = FindPortfolioPath()
var PortfolioPath string = PortfolioDir + "/portfolio.json"

func FindPortfolioPath() string {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println(err)
    }
    return homeDir + LINUX_PREFIX
}

func OpenPortfolio() *[]coin.Coin {
    // portfolio is a slice of Coins
    portfolio := []coin.Coin{}
    // check if portfolio exists
    fmt.Println(file.Exists(PortfolioPath))
    if !file.Exists(PortfolioPath) {
        CreatePortfolio()
    }
    // open portfolio location
	portfolioJson, err := ioutil.ReadFile(PortfolioPath)
    // if portfolio could not be found
	if err != nil {
		fmt.Println("File could not be read, try 'kripto init'")
		return nil
	}
    fmt.Println(portfolioJson)
    // unmarshal into portfolio var
	json.Unmarshal(portfolioJson, &portfolio)
    // return address of portfolio array
	return &portfolio
}

func CreatePortfolio() {
    // make directory
    os.MkdirAll(PortfolioDir, os.ModePerm)
    // create portfolio at PortfolioDir
    _, err := os.Create(PortfolioPath)
    if err != nil {
        fmt.Println("Cannot create new portfolio.")
        return
    }
    return
}

func SavePortfolio(portfolio *[]coin.Coin) {
	portfolioJson, err := json.MarshalIndent(*portfolio, "", "   ")
	if err != nil {
		fmt.Println("Error encoding as json.")
	}
	err = ioutil.WriteFile(PortfolioPath, portfolioJson, 0770)
	if err != nil {
		fmt.Println("Error writing to file.")
	}
	return
}
