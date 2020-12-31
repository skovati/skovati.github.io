package main

import (
    "fmt"
    // "github.com/skovati/kripto/coin"
    // "github.com/skovati/kripto/file"
    "github.com/skovati/kripto/portfolio"
)

func main() {
    folio := *portfolio.OpenPortfolio()
    fmt.Println(folio[0].ID)
}
