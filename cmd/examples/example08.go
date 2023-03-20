package main

import (
	"fmt"
)

// 構造体を定義
type pokemon struct {
	Id   int
	Name string
	Type string
}

func main() {
	// pokemon型の構造体を宣言
	var charizard pokemon

	// 構造体の各フィールドに値を代入
	charizard.Id = 6
	charizard.Name = "リザードン"
	charizard.Type = "ほのお"

	fmt.Printf("図鑑No:%d\n", charizard.Id)
	fmt.Println(charizard.Name)
	fmt.Printf("タイプ:%s\n", charizard.Type)
}
