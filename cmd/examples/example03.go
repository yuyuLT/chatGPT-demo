package main

import (
	"fmt"
)

func main() {
	var change int
	fmt.Print("お釣りを入力してください > ")
	fmt.Scan(&change)

	fmt.Println(change)
}
