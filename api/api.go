package api

import (
	"net/http"
    "os"
)


var CACHE_DIR string = os.GetUserCacheDir + "/kripto"

type APICoin []struct{
  ID string `json:"id"`
  Symbol string `json:"symbol"`
  Name string `json:"name"`
}

func getPrice(id string) (usdPrice, percentChange1H, percentChange24H, percentChange7D float64) {
	id = strings.ToLower(ticker_to_id(currency))
	url := "https://api.coingecko.com/api/v3/coins/" + currency + "?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]map[string]map[string]float64
	json.Unmarshal([]byte(body), &result)

	usdPrice = result["market_data"]["current_price"]["usd"]
	percentChange1H = result["market_data"]["price_change_percentage_1h_in_currency"]["usd"]
	percentChange24H = result["market_data"]["price_change_percentage_24h_in_currency"]["usd"]
	percentChange7D = result["market_data"]["price_change_percentage_7d_in_currency"]["usd"]

	return
}

func GetID(currency string) string {
    // if coin cache doesn't exist, run api cache function
    if _, err := os.Stat(CACHE_DIR + "/kripto_coins.json"); os.IsNotExist(err) {
        if !GetSupportedCoins() {
            fmt.Println("Error fetching results, check internet connection")
        }
    }

    // loop through, if there is a match between name or ticket, return id
    portfolioJson, err := ioutil.ReadFile("/tmp/coins.json")
    handle_err(err)

    var list APICoin
    json.Unmarshal(portfolioJson, &list)
    for _,c := range list {
      if c.Symbol == ticker {
        return c.ID
      }
    }
    return currency
}

func GetSupportedCoins() bool {
	url := "https://api.coingecko.com/api/v3/coins/list"
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
    }

	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
    }

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

    // make 
    err = os.MkdirAll(CACHE_DIR, os.ModePerm)
    if err != nil {
       fmt.Println(err)
    }

    // 
	err = ioutil.WriteFile(CACHE_DIR + "/kripto_coins.json", body, os.ModePerm)
	return err == nil
}
