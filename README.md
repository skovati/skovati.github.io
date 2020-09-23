# kripto
Simple cryptocurrency tracker written in Go.
Uses the [CoinGecko](https://www.coingecko.com/en/api) api to fetch real time prices and persistently stores the portfolio in json format.

### installation
```
git clone https://github.com/skovati/kripto
cd kripto
make && sudo make install
```

### usage
```
kripto init - Create a new portfolio
kripto add - Use 'kripto add <currency> <amount>' to add a new crypto
kripto remove - Use 'kripto remove <currency>' to remove
kripto update - Use 'kripto update <currency> <new amount>' to update holdings
kripto list - Lists current porfolio and holdings
kripto help - Show this message
```
