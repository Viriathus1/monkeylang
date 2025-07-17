package evaluator

import (
	"monkeylang/ast"
	"monkeylang/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
