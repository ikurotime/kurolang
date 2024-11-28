package parser

import (
	"fmt"
	"kuro/kurolang/internal/ast"
	"kuro/kurolang/internal/token"
)

const (
    ErrorUnexpectedToken = "unexpected token type %s"
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
   if p.LookAhead().Type == "OPERATOR" && p.LookAhead().Value == "=" {
       value, err :=  p.ParseAssignment()
       return value, nil, err
    }else{
        value,err := p.ParseExpression()
        return ast.Assignment{}, value, err
    }
}

func (p *Parser) ParseAssignment() (ast.Assignment, error) {
    identifier, err := p.Consume("IDENTIFIER")
	if err != nil {
		return ast.Assignment{}, err
	}

	_, err = p.Consume("OPERATOR")
	if err != nil {
		return ast.Assignment{}, err
	}

	value, err := p.ParseExpression()
    if err != nil {
        p.addError(err.Error())
        return  ast.Assignment{}, err
    }
	return ast.Assignment{Identifier: identifier.Value, Value: value}, nil
}

func (p *Parser) ParseExpression() (ast.Node, error) {
	left, err := p.ParseTerm()
    if err != nil {
        p.addError(err.Error())
        return  ast.Number{Value: ""}, err
    }
	for p.CurrentToken().Value == "+" || p.CurrentToken().Value == "-" {
		operator, err := p.Consume("OPERATOR")
        if err != nil {
            p.addError(err.Error())
            return  ast.Number{Value: ""}, err
        }
		right, err := p.ParseTerm()
        if err != nil {
            p.addError(err.Error())
            return  ast.Number{Value: ""}, err
        }
		left = ast.BinaryOperation{Operator: operator.Value, Left: left, Right: right}
	}

	return left, nil
}

func (p *Parser) ParseTerm() (ast.Node, error) {
	left, err := p.ParseFactor()
    if err != nil {
        p.addError(err.Error())
        return  ast.Number{Value: ""}, err
    }

	for p.CurrentToken().Value == "*" || p.CurrentToken().Value == "/" { 
		operator, err := p.Consume("OPERATOR")
        if err != nil {
            p.addError(err.Error())
        }   
		right, err := p.ParseFactor()
        if err != nil {
            p.addError(err.Error())
            return  ast.Number{Value: ""}, err
        }
		left = ast.BinaryOperation{Operator: operator.Value, Left: left, Right: right}
	}

	return left, nil
}

func (p *Parser) ParseFactor() (ast.Node,error) {
	token := p.CurrentToken()
    if token.Type == "NUMBER" {
        number, err :=  p.Consume("NUMBER")
        if err != nil {
            return ast.Number{Value: ""}, err
        }
        return ast.Number{Value: number.Value}, nil
	}else if token.Type == "IDENTIFIER" {
        identifier, err :=  p.Consume("IDENTIFIER")
        if err != nil {
            return ast.Identifier{Value: ""}, err
        }
        return ast.Identifier{Value: identifier.Value}, nil
    } else if token.Value == "(" {
        p.Consume("OPERATOR")
        expr, err := p.ParseExpression()
        if err != nil {
            return ast.Number{Value: ""}, err
        }
        p.Consume("OPERATOR")
        return expr, nil
    }
    return ast.Number{Value: ""}, fmt.Errorf(ErrorUnexpectedToken, token.Type)
}
