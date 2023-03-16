package main

import (
	"fmt"
)

// 構造体を定義
type Pokemon struct {
	Id   int
	Name string
	Type string
}

func main() {
	// Pokemon型の構造体を宣言
	var pokemon Pokemon

	// 構造体の各フィールドに値を代入
	pokemon.Id = 6
	pokemon.Name = "リザードン"
	pokemon.Type = "ほのお"

	fmt.Printf("図鑑No:%d\n", pokemon.Id)
	fmt.Println(pokemon.Name)
	fmt.Printf("タイプ:%s\n", pokemon.Type)
}
