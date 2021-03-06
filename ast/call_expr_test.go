package ast

import (
	"testing"
)

func TestCallExpr(t *testing.T) {
	nodes := map[string]Node{
		`0x7f9bf3033240 <col:11, col:25> 'int'`: &CallExpr{
			Address:  "0x7f9bf3033240",
			Position: "col:11, col:25",
			Type:     "int",
			Children: []Node{},
		},
		`0x7f9bf3035c20 <line:7:4, col:64> 'int'`: &CallExpr{
			Address:  "0x7f9bf3035c20",
			Position: "line:7:4, col:64",
			Type:     "int",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
