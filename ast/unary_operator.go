package ast

import (
	"fmt"

	"github.com/elliotchance/c2go/program"
	"github.com/elliotchance/c2go/types"
	"github.com/elliotchance/c2go/util"
)

type UnaryOperator struct {
	Address  string
	Position string
	Type     string
	IsLvalue bool
	IsPrefix bool
	Operator string
	Children []Node
}

func parseUnaryOperator(line string) *UnaryOperator {
	groups := groupsFromRegex(
		`<(?P<position>.*)>
		 '(?P<type>.*?)'
		(?P<lvalue> lvalue)?
		(?P<prefix> prefix)?
		(?P<postfix> postfix)?
		 '(?P<operator>.*?)'`,
		line,
	)

	return &UnaryOperator{
		Address:  groups["address"],
		Position: groups["position"],
		Type:     groups["type"],
		IsLvalue: len(groups["lvalue"]) > 0,
		IsPrefix: len(groups["prefix"]) > 0,
		Operator: groups["operator"],
		Children: []Node{},
	}
}

func (n *UnaryOperator) render(program *program.Program) (string, string) {
	operator := n.Operator
	expr, exprType := renderExpression(program, n.Children[0])

	if operator == "!" {
		if exprType == "bool" || exprType == "_Bool" {
			return fmt.Sprintf("!(%s)", expr), exprType
		}

		program.AddImport("github.com/elliotchance/c2go/noarch")

		t := types.ResolveType(program, exprType)
		functionName := fmt.Sprintf("noarch.Not%s", util.Ucfirst(t))
		return fmt.Sprintf("%s(%s)", functionName, expr), exprType
	}

	if operator == "*" {
		if exprType == "const char *" {
			return fmt.Sprintf("%s[0]", expr), "char"
		}

		t, err := types.GetDereferenceType(exprType)
		if err != nil {
			panic(err)
		}

		return fmt.Sprintf("*%s", expr), t
	}

	if operator == "++" {
		return fmt.Sprintf("%s += 1", expr), exprType
	}

	if operator == "~" {
		operator = "^"
	}

	if operator == "&" {
		exprType += " *"
	}

	return fmt.Sprintf("%s%s", operator, expr), exprType
}

func (n *UnaryOperator) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
