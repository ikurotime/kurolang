package ast

import "fmt"

type Node interface {
	String() string
}

type Assignment struct {
	Identifier string
	Value      Node
}

func (a Assignment) String() string {
	return fmt.Sprintf("Assignment(%s = %s)", a.Identifier, a.Value.String())
}

type BinaryOperation struct {
	Operator string
	Left     Node
	Right    Node
}

func (b BinaryOperation) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator, b.Right.String())
}

type Number struct {
	Value string
}

func (n Number) String() string {
	return n.Value
}

type Variable struct {
	Value string
}

func (v Variable) String() string {
	return fmt.Sprintf("var %s", v.Value)
}

type Identifier struct {
	Value string
}

func (i Identifier) String() string {
	return i.Value
}

type Program struct {
	Statements []Node
}

func (p Program) String() string {
	result := "Program{\n"
	for _, stmt := range p.Statements {
		result += "  " + stmt.String() + "\n"
	}
	result += "}"
	return result
}
