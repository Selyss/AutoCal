package parser

import (
	"testing"

	"github.com/Selyss/AutoCal/ast"
	"github.com/Selyss/AutoCal/lexer"
)

func TestAddStatements(t *testing.T) {
	input := `
ADD "Event".
ADD "Another Event".
ADD "Third Event".
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"Event"},
		{"Another Event"},
		{"Third Event"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAddStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testAddStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "ADD" {
		t.Errorf("s.TokenLiteral not 'ADD'. got=%q", s.TokenLiteral())
		return false
	}
	addStmt, ok := s.(*ast.AddStatement)
	if !ok {
		t.Errorf("s not *ast.AddStatement. got=%T", s)
		return false
	}
	if addStmt.Name.Value != name {
		t.Errorf("addStmt.Name.Value not '%s'. got=%s", name, addStmt.Name.Value)
		return false
	}
	if addStmt.Name.TokenLiteral() != name {
		t.Errorf("addStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, addStmt.Name.TokenLiteral())
		return false
	}
	return true
}
