package lexer

import (
	"kuro/kurolang/internal/token"
	"testing"
)

type test struct {
	expectedType    string
	expectedLiteral string
}

func RunTests(tokens []token.Token, tests []test, t *testing.T) {
	for i, tt := range tests {
		if tokens[i].Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tokens[i].Type)
		}

		if tokens[i].Value != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tokens[i].Value)
		}
	}
}

func TestTokens(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	let result = x + y;
	`

	tests := []test{
		{token.LET, "let"},
		{token.IDENTIFIER, "x"},
		{token.OPERATOR, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "y"},
		{token.OPERATOR, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.OPERATOR, "="},
		{token.IDENTIFIER, "x"},
		{token.OPERATOR, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
	}

	tokens, err := Tokenize(input)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	RunTests(tokens, tests, t)
}

// We can test lexer (and more inputs) like this intead of running main.go
func TestVarAssignment(t *testing.T) {
	input := `let x = 5 + 5;`

	tests := []test{
		{token.LET, "let"},
		{token.IDENTIFIER, "x"},
		{token.OPERATOR, "="},
		{token.NUMBER, "5"},
		{token.OPERATOR, "+"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
	}

	tokens, err := Tokenize(input)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	RunTests(tokens, tests, t)
}
