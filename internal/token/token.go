package token

const (
	// KEYWORDS
	CONST              = "const"
	LET                = "let"
	IDENTIFIER         = "[a-zA-Z]+"
	NUMBER             = "[0-9]+"
	OPERATOR           = "[+\\-*/=()]"
	WHITESPACE         = "\\s+"
	SEMICOLON          = ";"
	MUST_COMPILE_REGEX = "[a-zA-Z]+|[0-9]+|[\\+\\-\\=\\;]"
)

var TOKENS = map[string]string{
	IDENTIFIER: "IDENTIFIER",
	NUMBER:     "NUMBER",
	OPERATOR:   "OPERATOR",
	SEMICOLON:  "SEMICOLON",
	WHITESPACE: "WHITESPACE",
}

// Token struct
type Token struct {
	Type  string
	Value string
}
