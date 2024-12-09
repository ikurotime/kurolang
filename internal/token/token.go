package token

const (
	// KEYWORDS
	CONST              = "const"
	LET                = "let"
	IDENTIFIER         = "[a-zA-Z]+"
	NUMBER             = "[0-9]+"
	OPERATOR           = "[+\\-*/=]"
	COMPARISON         = "[><=!]=|[><]"
	EQUALS             = "=="
	WHITESPACE         = "\\s+"
	SEMICOLON          = ";"
	MUST_COMPILE_REGEX = "[a-zA-Z]+|[0-9]+|[><=!]=|[><]|[+\\-*/=;{}]"
	IF                 = "if"
	ELSE               = "else"
	ELSEIF             = "else if"
	TRUE               = "true"
	FALSE              = "false"
	LBRACKET           = "{"
	RBRACKET           = "}"
)

// Token struct
type Token struct {
	Type  string
	Value string
}
