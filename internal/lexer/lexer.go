package lexer

import (
	"fmt"
	"kuro/kurolang/internal/token"
	"regexp"
)

type Lexer struct {
	Input    string
	Position int
}

const (
	ErrorInvalidToken = "invalid token: %s"
)

func SplitStringWithRegex(input string) []string {
	var result []string
	re := regexp.MustCompile(token.MUST_COMPILE_REGEX)
	result = re.FindAllString(input, -1)
	return result
}
func SplitString(input string) []string {
	var result []string
	start := 0
	end := 0

	for end < len(input) {
		if input[end] == ' ' {
			word := input[start:end]
			result = append(result, word)
			start = end + 1
		}
		end++
	}
	if start < end {
		word := input[start:end]
		result = append(result, word)
	}
	return result
}

// Tokenize the input
func Tokenize(input string) ([]token.Token, error) {
	var Tokens []token.Token

	inputString := SplitStringWithRegex(input)
	fmt.Printf("input tokens: %s\n", inputString)

	identifier, err := regexp.Compile(token.IDENTIFIER)
	if err != nil {
		return nil, err
	}
	operator, err := regexp.Compile(token.OPERATOR)
	if err != nil {
		return nil, err
	}
	number, err := regexp.Compile(token.NUMBER)
	if err != nil {
		return nil, err
	}

	for _, word := range inputString {
		fmt.Printf("Token: %s\n", word)
		if word == token.LET {
			Tokens = append(Tokens, NewToken(token.LET, word))
		} else if identifier.MatchString(word) {
			Tokens = append(Tokens, NewToken(token.IDENTIFIER, word))
		} else if operator.MatchString(word) {
			Tokens = append(Tokens, NewToken(token.OPERATOR, word))
		} else if number.MatchString(word) {
			Tokens = append(Tokens, NewToken(token.NUMBER, word))
		} else if word == token.SEMICOLON {
			Tokens = append(Tokens, NewToken(token.SEMICOLON, word))
		} else {
			return nil, fmt.Errorf(ErrorInvalidToken, word)
		}
	}

	return Tokens, nil
}

func NewToken(tokenType string, tokenValue string) token.Token {
	return token.Token{Type: tokenType, Value: tokenValue}
}
