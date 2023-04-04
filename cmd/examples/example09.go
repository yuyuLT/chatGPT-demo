package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 構造体を定義
type pokemon struct {
	Id   int
	Name string
	Type string
}

func main() {
	// jsonファイルの読み込み
	pokeData, err := os.ReadFile("pokemon.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// jsonを構造体に格納
	var charmander pokemon
	err = json.Unmarshal(pokeData, &charmander)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("図鑑No:%d\n", charmander.Id)
	fmt.Println(charmander.Name)
	fmt.Printf("タイプ:%s\n", charmander.Type)
}
