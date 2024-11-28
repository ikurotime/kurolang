package token

import (
	"fmt"
	"regexp"
)

const (
    // KEYWORDS
    VAR = "var"
    IDENTIFIER =  "[a-zA-Z]+"
    NUMBER = "[0-9]+"
    OPERATOR = "[+\\-*/=()]"
    WHITESPACE = "\\s+"
    MUST_COMPILE_REGEX = "[a-zA-Z]+|[0-9]+|[\\+\\-\\=\\;]"  
)

// String with regex patterns
var TOKENS = map[string]string{
	VAR: "VAR",
	IDENTIFIER: "IDENTIFIER",
	NUMBER: "NUMBER",
    OPERATOR : "OPERATOR",
	WHITESPACE: "WHITESPACE",
}


// Token struct
type Token struct {
	Type  string
	Value string
    Position int
}

// Token constructor
func NewToken(tokenType string, tokenValue string, tokenPosition int) Token {
    return Token{Type: tokenType, Value: tokenValue , Position: tokenPosition}
}

func SplitStringWithRegex(input string) []string {
	var result []string
	re := regexp.MustCompile(MUST_COMPILE_REGEX)
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
func Tokenize(input string) ([]Token, error) {
	var Tokens []Token
    position := 0

	inputString := SplitStringWithRegex(input)
    fmt.Printf("input tokens: %s\n",inputString)
	for _, token := range inputString {
        position++
		switch token {
		case VAR:
			Tokens = append(Tokens, NewToken(TOKENS[VAR], token, position))
			continue
		default:
			identifier, err := regexp.MatchString(IDENTIFIER, token)
			if err != nil {
				return nil, err
			}
			if identifier {
				Tokens = append(Tokens, NewToken(TOKENS[IDENTIFIER], token, position))
			}
            operator, err := regexp.MatchString(OPERATOR, token)
            if err != nil {
                return nil, err
            }
            if operator {
                Tokens = append(Tokens, NewToken(TOKENS[OPERATOR], token, position))
            }
			number, err := regexp.MatchString(NUMBER, token)
			if err != nil {
				return nil, err
			}
			if number {
				Tokens = append(Tokens, NewToken(TOKENS[NUMBER], token, position))
			}
		}
	}

	return Tokens, nil
}
