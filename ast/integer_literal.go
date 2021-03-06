package ast

import (
	"strconv"

	"github.com/elliotchance/c2go/program"
	"github.com/elliotchance/c2go/util"
)

type IntegerLiteral struct {
	Address  string
	Position string
	Type     string
	Value    int
	Children []Node
}

func parseIntegerLiteral(line string) *IntegerLiteral {
	groups := groupsFromRegex(
		"<(?P<position>.*)> '(?P<type>.*?)' (?P<value>\\d+)",
		line,
	)

	return &IntegerLiteral{
		Address:  groups["address"],
		Position: groups["position"],
		Type:     groups["type"],
		Value:    util.Atoi(groups["value"]),
		Children: []Node{},
	}
}

func (n *IntegerLiteral) render(program *program.Program) (string, string) {
	literal := n.Value

	// FIXME
	//if str(literal)[-1] == 'L':
	//    literal = '%s(%s)' % (resolveType('long'), literal[:-1])

	return strconv.FormatInt(int64(literal), 10), "int"
}

func (n *IntegerLiteral) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
