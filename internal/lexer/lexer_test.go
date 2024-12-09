package lexer

import (
	"fmt"
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
	const y = 10;
	let result = x + y;
	`
	tests := []test{
		{token.LET, "let"},
		{token.IDENTIFIER, "x"},
		{token.OPERATOR, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.CONST, "const"},
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

func TestLetAssignment(t *testing.T) {
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

func TestConstAssignment(t *testing.T) {
	input := `const x = 5 + 5;`

	tests := []test{
		{token.CONST, "const"},
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

func TestMultiAssignment(t *testing.T) {
	input := `const x = 5 + 5;`

	tests := []test{
		{token.CONST, "const"},
		{token.IDENTIFIER, "x"},
		{token.OPERATOR, "="},
		{token.NUMBER, "5"},
		{token.OPERATOR, "+"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
	}

	tokens, err := Tokenize(input)

	fmt.Printf("Tokens: %v\n", tokens)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	RunTests(tokens, tests, t)
}

func TestControlStructures(t *testing.T) {
	input := `const x = 5 + 5;
	if x == 10 {
		let y = 10;
};
	`

	tests := []test{
		{token.CONST, "const"},
		{token.IDENTIFIER, "x"},
		{token.OPERATOR, "="},
		{token.NUMBER, "5"},
		{token.OPERATOR, "+"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.IDENTIFIER, "x"},
		{token.EQUALS, "=="},
		{token.NUMBER, "10"},
		{token.LBRACKET, "{"},
		{token.LET, "let"},
		{token.IDENTIFIER, "y"},
		{token.OPERATOR, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.SEMICOLON, ";"},
	}

	tokens, err := Tokenize(input)

	fmt.Printf("Tokens: %v\n", tokens)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	RunTests(tokens, tests, t)
}

func TestBooleans(t *testing.T) {
	input := `const x = true;
	const y = false;
	`
	tests := []test{
		{token.CONST, "const"},
		{token.IDENTIFIER, "x"},
		{token.OPERATOR, "="},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.CONST, "const"},
		{token.IDENTIFIER, "y"},
		{token.OPERATOR, "="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
	}

	tokens, err := Tokenize(input)

	fmt.Printf("Tokens: %v\n", tokens)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	RunTests(tokens, tests, t)
}

func TestComparison(t *testing.T) {
	input := `x > 5; y < 10; z == 5; a >= 5; b <= 10; c != 5;`
	tests := []test{
		{token.IDENTIFIER, "x"},
		{token.COMPARISON, ">"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "y"},
		{token.COMPARISON, "<"},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "z"},
		{token.COMPARISON, "=="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "a"},
		{token.COMPARISON, ">="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "b"},
		{token.COMPARISON, "<="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "c"},
		{token.COMPARISON, "!="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
	}

	tokens, err := Tokenize(input)

	fmt.Printf("Tokens: %v\n", tokens)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	RunTests(tokens, tests, t)
}
