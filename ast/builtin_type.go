package ast

import (
	"github.com/elliotchance/c2go/program"
)

type BuiltinType struct {
	Address  string
	Type     string
	Children []Node
}

func parseBuiltinType(line string) *BuiltinType {
	groups := groupsFromRegex(
		"'(?P<type>.*?)'",
		line,
	)

	return &BuiltinType{
		Address:  groups["address"],
		Type:     groups["type"],
		Children: []Node{},
	}
}

func (n *BuiltinType) render(program *program.Program) (string, string) {
	return "", ""
}

func (n *BuiltinType) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
