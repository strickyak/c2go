package ast

import (
	"testing"
)

func TestRecordType(t *testing.T) {
	nodes := map[string]Node{
		`0x7fd3ab84dda0 'struct _opaque_pthread_condattr_t'`: &RecordType{
			Address:  "0x7fd3ab84dda0",
			Type:     "struct _opaque_pthread_condattr_t",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
