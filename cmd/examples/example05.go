package main

import (
	"fmt"
)

func main() {
	var change int
	fmt.Print("お釣りを入力してください > ")
	fmt.Scan(&change)

	currency := []int{5000, 1000, 500, 100, 50, 10, 5, 1}

	for _, v := range currency {
		count := change / v
		change -= v * count

		fmt.Println(v, count)
	}
}
