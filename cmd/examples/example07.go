package main

import (
	"fmt"
)

func main() {
	var change int
	fmt.Print("お釣りを入力してください > ")

	fmt.Scan(&change)

	currency := []int{5000, 1000, 500, 100, 50, 10, 5, 1}

	fmt.Println(currency)

	countNumber(change, currency)
}

func countNumber(change int, currency []int) {
	for _, v := range currency {
		count := change / v
		change -= v * count

		if count != 0 {
			fmt.Println(v, count)
		}
	}
}
