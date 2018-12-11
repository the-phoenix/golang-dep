package main

import (
	"fmt"
	humanize "github.com/dustin/go-humanize"
)

func main() {
	fmt.Println("hello world")
	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982)) // That file is 83 MB.
	fmt.Printf("You're my %s best friend.\n", humanize.Ordinal(193)) // You are my 193rd best friend.
	fmt.Printf("You owe $%s.\n", humanize.Comma(6582491)) // You owe $6,582,491.
}