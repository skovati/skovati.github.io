# kripto
Simple cryptocurrency tracker written in Go.
Uses the [CoinGecko](https://www.coingecko.com/en/api) api to fetch real time prices and persistently stores the portfolio in json format.

### installation
Unix
```
git clone https://github.com/skovati/kripto
cd kripto
sudo make install
```
Gophers can build from source with 
```
make && sudo make install
```
Windows users can run the exe found in the bin/ folder
```
& bin/kripto.exe
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

