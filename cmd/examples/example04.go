package main

import (
	"fmt"
)

func main() {
	var change int
	fmt.Print("お釣りを入力してください > ")
	fmt.Scan(&change)

	fmt.Println(change)

	currency := []int{5000, 1000, 500, 100, 50, 10, 5, 1}

	fmt.Println(currency)
}
