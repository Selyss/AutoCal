package lexer

import (
	"testing"

	"github.com/Selyss/AutoCal/token"
)

func TestNextToken(t *testing.T) {
	// 	input := `
	//   IF DAY ODD, ADD "quite an odd day!", TIME 12:30pm.
	// `
	// input := `IF DAY ODD, ADD "quite an odd day!".`
	input := `
	IF DAY ODD, ADD "quite an odd day!".

	IF DAY EVEN, ADD (
		TITLE "Event!",
		DESC "This is an event!",
		COLOR RED
	).
	ELSE, ADD (
		TITLE "Other Event!",
		DESC "This is another event!",
		COLOR BLUE
	).`

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
		{token.IF, "IF"},
		{token.DAY, "DAY"},
		{token.EVEN, "EVEN"},
		{token.COMMA, ","},
		{token.ADD, "ADD"},
		{token.LPAREN, "("},
		{token.TITLE, "TITLE"},
		{token.STRING, "Event!"},
		{token.COMMA, ","},
		{token.DESC, "DESC"},
		{token.STRING, "This is an event!"},
		{token.COMMA, ","},

		// TODO: rethink how colors are handled
		{token.COLOR, "COLOR"},
		{token.RED, "RED"},
		{token.RPAREN, ")"},
		{token.PERIOD, "."},
		{token.ELSE, "ELSE"},
		{token.COMMA, ","},
		{token.ADD, "ADD"},
		{token.LPAREN, "("},
		{token.TITLE, "TITLE"},
		{token.STRING, "Other Event!"},
		{token.COMMA, ","},
		{token.DESC, "DESC"},
		{token.STRING, "This is another event!"},
		{token.COMMA, ","},
		{token.COLOR, "COLOR"},
		{token.BLUE, "BLUE"},
		{token.RPAREN, ")"},
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
