package token

import (
	"regexp"
)

// String with regex patterns
var TOKENS = map[string]string{
	// KEYWORDS
	"var": "VAR",
	// IDENTIFIERS
	"[a-zA-Z]+": "IDENTIFIER",
	// LITERALS
	"[0-9]+": "NUMBER",
	// OPERATORS
	"+": "PLUS",
	"-": "MINUS",
	"=": "ASSIGN",
	";": "SEMICOLON",
	// SYMBOLS
	" ": "WHITESPACE",
	//TODO- ADD PUNCTUATION
	//"(": "LPAREN",
	//")": "RPAREN",
	//"{": "LBRACE",
	//"}": "RBRACE",
	//"[": "LBRACKET",
	//"]": "RBRACKET",
}

// Token struct
type Token struct {
	Type  string
	Value string
}

// Token constructor
func NewToken(tokenType string, tokenValue string) Token {
	return Token{Type: tokenType, Value: tokenValue}
}

func SplitStringWithRegex(input string) []string {
	var result []string
	re := regexp.MustCompile(`[a-zA-Z]+|[0-9]+|[\+\-\=\;]`)
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
	inputString := SplitStringWithRegex(input)
	//   fmt.Printf("input tokens: %s\n",inputString)
	for _, token := range inputString {
		switch token {
		case "var":
			Tokens = append(Tokens, NewToken(TOKENS["var"], token))
			continue
		case "+":
			Tokens = append(Tokens, NewToken(TOKENS["+"], token))
			continue
		case "-":
			Tokens = append(Tokens, NewToken(TOKENS["-"], token))
			continue
		case "=":
			Tokens = append(Tokens, NewToken(TOKENS["="], token))
			continue
		case ";":
			Tokens = append(Tokens, NewToken(TOKENS[";"], token))
			continue
		default:
			identifier, err := regexp.MatchString("[a-zA-Z]+", token)
			if err != nil {
				return nil, err
			}
			if identifier {
				Tokens = append(Tokens, NewToken(TOKENS["[a-zA-Z]+"], token))
			}
			integer, err := regexp.MatchString("[0-9]+", token)
			if err != nil {
				return nil, err
			}
			if integer {
				Tokens = append(Tokens, NewToken(TOKENS["[0-9]+"], token))
			}
		}
	}

	return Tokens, nil
}
