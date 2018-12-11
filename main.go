package main

import (
	"fmt"
	"math/big"

	accounting "github.com/leekchan/accounting"
	humanize "github.com/dustin/go-humanize"
)

func main() {
	fmt.Println("hello world")
	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982)) // That file is 83 MB.
	fmt.Printf("You're my %s best friend.\n", humanize.Ordinal(193)) // You are my 193rd best friend.
	fmt.Printf("You owe $%s.\n", humanize.Comma(6582491)) // You owe $6,582,491.

	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))                       // "$123,456,789.21"
	fmt.Println(ac.FormatMoney(12345678))                               // "$12,345,678.00"
	fmt.Println(ac.FormatMoney(big.NewRat(77777777, 3)))                // "$25,925,925.67"
	fmt.Println(ac.FormatMoney(big.NewRat(-77777777, 3)))               // "-$25,925,925.67"
	fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(123456789.213123))) // "$123,456,789.21"
}