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

	equals, err := regexp.Compile(token.EQUALS)
	if err != nil {
		return nil, err
	}
	comparison, err := regexp.Compile(token.COMPARISON)
	if err != nil {
		return nil, err
	}
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
		switch {
		case word == token.CONST:
			Tokens = append(Tokens, NewToken(token.CONST, word))
		case word == token.LET:
			Tokens = append(Tokens, NewToken(token.LET, word))
		case word == token.IF:
			Tokens = append(Tokens, NewToken(token.IF, word))
		case word == token.ELSE:
			Tokens = append(Tokens, NewToken(token.ELSE, word))
		case word == token.ELSEIF:
			Tokens = append(Tokens, NewToken(token.ELSEIF, word))
		case word == token.LBRACKET:
			Tokens = append(Tokens, NewToken(token.LBRACKET, word))
		case word == token.RBRACKET:
			Tokens = append(Tokens, NewToken(token.RBRACKET, word))
		case word == token.TRUE:
			Tokens = append(Tokens, NewToken(token.TRUE, word))
		case word == token.FALSE:
			Tokens = append(Tokens, NewToken(token.FALSE, word))
		case word == token.SEMICOLON:
			Tokens = append(Tokens, NewToken(token.SEMICOLON, word))
		case word == token.TRUE:
			Tokens = append(Tokens, NewToken(token.TRUE, word))
		case word == token.FALSE:
			Tokens = append(Tokens, NewToken(token.FALSE, word))
		case comparison.MatchString(word):
			Tokens = append(Tokens, NewToken(token.COMPARISON, word))
		case identifier.MatchString(word):
			Tokens = append(Tokens, NewToken(token.IDENTIFIER, word))
		case equals.MatchString(word):
			Tokens = append(Tokens, NewToken(token.EQUALS, word))
		case operator.MatchString(word):
			Tokens = append(Tokens, NewToken(token.OPERATOR, word))
		case number.MatchString(word):
			Tokens = append(Tokens, NewToken(token.NUMBER, word))
		default:
			return nil, fmt.Errorf(ErrorInvalidToken, word)
		}
	}

	return Tokens, nil
}

func NewToken(tokenType string, tokenValue string) token.Token {
	return token.Token{Type: tokenType, Value: tokenValue}
}
