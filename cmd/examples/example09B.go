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
	Type []string
}

func main() {
	// jsonファイルの読み込み
	pokeData, err := os.ReadFile("pokemon2.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// jsonを構造体に格納
	var charizard pokemon
	err = json.Unmarshal(pokeData, &charizard)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("図鑑No:%d\n", charizard.Id)
	fmt.Println(charizard.Name)
	fmt.Print("タイプ:")
	for _, v := range charizard.Type {
		fmt.Printf("%s ", v)
	}
}
