package main

import (
	"fmt"
	"os"
	"tiny-basic/src/tokenizer"
)

func main() {
	fileName := "example.tb"
	sourceCode, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	tokens, err := tokenizer.Tokenize(string(sourceCode))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range tokens {
		fmt.Printf("Token(Type: %s, Value: %s)\n", token.Type, token.Value)
	}
}
