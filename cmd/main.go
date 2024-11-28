package main

import (
	"fmt"
	"kuro/kurolang/internal/parser"
	"kuro/kurolang/internal/token"
)

func main() {
	var x = "x = 5 + 5;"

	// Tokenize the input
	tokens, err := token.Tokenize(x)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// Print the tokens
	for _, token := range tokens {
		fmt.Printf("Token: {%s,%s}\n", token.Type, token.Value)
	}

	parser := parser.Parser{Tokens: tokens, Position: 0}
	parser.PrintAST()

}
