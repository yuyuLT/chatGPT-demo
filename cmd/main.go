package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}
type Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {
	apiKey := os.Getenv("API_KEY")
	apiURL := "https://api.openai.com/v1/engines/text-davinci-003/completions"
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your question: ")
	prompt, _ := reader.ReadString('\n')
	prompt = strings.TrimSpace(prompt)
	maxTokens := 300
	temperature := float32(0.5)
	requestBody, err := json.Marshal(Request{prompt, maxTokens, temperature})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get response, status code: %d\n", resp.StatusCode)
		return
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	var sb strings.Builder
	for _, choice := range response.Choices {
		sb.WriteString(choice.Text)
	}
	fmt.Println(sb.String())
}
