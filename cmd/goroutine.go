package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, concurrency!")

	go printMessage()
	go printMessage2()

	//3秒待つ
	time.Sleep(time.Second * 3)

	fmt.Println("finish program!")
}

// 1秒待った後、メッセージを出力する関数
func printMessage() {
	time.Sleep(time.Second * 1)
	fmt.Println("this is message 1")
}

// 1秒待った後、メッセージを出力する関数2
func printMessage2() {
	time.Sleep(time.Second * 1)
	fmt.Println("this is message 2")
}
