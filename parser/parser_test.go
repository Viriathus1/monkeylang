package parser

import (
	"monkeylang/ast"
	"monkeylang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn't contain 3 statements, got=%d", len(program.Statements))
	}

	expectedIdentifiers := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range expectedIdentifiers {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got=%q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not %q, got=%s", letStatement.Name.Value, name)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.TokenLiteral not %q, got=%s", letStatement.Name.TokenLiteral(), name)
		return false
	}

	return true
}
