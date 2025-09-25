package evaluator

import (
	"monkeylang/object"
)

var builtins = map[string]*object.BuiltIn{
	"len":   object.GetBuiltInByName("len"),
	"first": object.GetBuiltInByName("first"),
	"last":  object.GetBuiltInByName("last"),
	"rest":  object.GetBuiltInByName("rest"),
	"push":  object.GetBuiltInByName("push"),
	"puts":  object.GetBuiltInByName("puts"),
}
