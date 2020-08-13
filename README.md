# kripto
Simple cryptocurrency tracker written in Go.
Uses the [CoinGecko](https://www.coingecko.com/en/api) api to fetch real time prices and persistently stores the portfolio in json format.

### usage
```
git clone https://github.com/skovati/kripto
cd kripto/bin
```
```
kripto init - create a new portfolio
kripto add - Use 'kripto add <currency> <amount>' to add a new crypto
kripto remove - Use 'kripto remove <currency>' to remove
kripto update - Use 'kripto update <currency>' <new amount> to update holdings
kripto list - lists current porfolio and holdings
kripto help - show this message
```
