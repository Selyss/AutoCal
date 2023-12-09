package lexer

import (
	"testing"

	"github.com/Selyss/AutoCal/token"
)

func TestNextToken(t *testing.T) {
	// 	input := `
	//   IF DAY ODD, ADD "quite an odd day!", TIME 12:30pm.
	// `
	input := `
  IF DAY ODD, ADD "quite an odd day!".
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "IF"},
		{token.DAY, "DAY"},
		{token.ODD, "ODD"},
		{token.COMMA, ","},
		{token.ADD, "ADD"},
		{token.STRING, "quite an odd day!"},
		{token.PERIOD, "."},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
