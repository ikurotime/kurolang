package parser

import (
	"fmt"
	"kuro/kurolang/internal/ast"
	"kuro/kurolang/internal/token"
)

type Parser struct {
	Tokens   []token.Token
	Position int
	Errors   []string
}

func (p Parser) CurrentToken() token.Token {
	if p.Position >= len(p.Tokens) {
		return token.Token{Type: "EOF", Value: ""}
	}
	return p.Tokens[p.Position]
}

func (p *Parser) addError(msg string) {
	p.Errors = append(p.Errors, msg)
}

func (p *Parser) Consume(expected_type string) (token.Token, error) {
	token := p.CurrentToken()
	if token.Type == expected_type {
		p.Position++
		return token, nil
	}
	err := fmt.Errorf("expected token type %s, got %s", expected_type, token.Type)
	p.addError(err.Error())
	return token, err
}

func (p *Parser) PrintAST() {
	assignment, expression, err := p.Parse()
	if err != nil {
		fmt.Printf("Parsing errors:\n")
		for _, err := range p.Errors {
			fmt.Printf("  - %s\n", err)
		}
		return
	}
	fmt.Printf("Assignment: %v\n", assignment)
	fmt.Printf("Expression: %v\n", expression)
}

func (p *Parser) LookAhead() token.Token {
	if p.Position+1 >= len(p.Tokens) {
		return token.Token{Type: "EOF", Value: ""}
	}
	return p.Tokens[p.Position+1]
}

func (p *Parser) Parse() (ast.Assignment, ast.Node, error) {
	if len(p.Tokens) == 0 {
		return ast.Assignment{}, nil, fmt.Errorf("no tokens to parse")
	}
	return p.ParseStatement()
}

func (p *Parser) ParseStatement() (ast.Assignment, ast.Node, error) {
	if p.CurrentToken().Type == "VAR" {
		assignment, err := p.ParseAssignment()
		return assignment, ast.Number{Value: ""}, err
	}
	expr := p.ParseExpression()
	return ast.Assignment{}, expr, nil
}

func (p *Parser) ParseAssignment() (ast.Assignment, error) {
	_, err := p.Consume("VAR")
	if err != nil {
		return ast.Assignment{}, err
	}

	identToken, err := p.Consume("IDENTIFIER")
	if err != nil {
		return ast.Assignment{}, err
	}

	_, err = p.Consume("ASSIGN")
	if err != nil {
		return ast.Assignment{}, err
	}

	value := p.ParseExpression()
	return ast.Assignment{Identifier: identToken.Value, Value: value}, nil
}

func (p *Parser) ParseExpression() ast.Node {
	left := p.ParseTerm()

	for p.CurrentToken().Type == "PLUS" || p.CurrentToken().Type == "MINUS" {
		operator := p.CurrentToken().Value
		if p.CurrentToken().Type == "PLUS" {
			p.Consume("PLUS")
		} else {
			p.Consume("MINUS")
		}
		right := p.ParseTerm()
		left = ast.BinaryOperation{Operator: operator, Left: left, Right: right}
	}

	return left
}

func (p *Parser) ParseTerm() ast.Node {
	left := p.ParseFactor()

	for p.CurrentToken().Type == "MULTIPLY" || p.CurrentToken().Type == "DIVIDE" {
		operator := p.CurrentToken().Value
		if p.CurrentToken().Type == "MULTIPLY" {
			p.Consume("MULTIPLY")
		} else {
			p.Consume("DIVIDE")
		}
		right := p.ParseFactor()
		left = ast.BinaryOperation{Operator: operator, Left: left, Right: right}
	}

	return left
}

func (p *Parser) ParseFactor() ast.Node {
	token := p.CurrentToken()
	switch token.Type {
	case "NUMBER":
		p.Consume("NUMBER")
		return ast.Number{Value: token.Value}
	case "IDENTIFIER":
		p.Consume("IDENTIFIER")
		return ast.Identifier{Value: token.Value}
	case "LPAREN":
		p.Consume("LPAREN")
		expr := p.ParseExpression()
		p.Consume("RPAREN")
		return expr
	default:
		p.addError(fmt.Sprintf("unexpected token %s", token.Type))
		return ast.Number{Value: "0"}
	}
}
