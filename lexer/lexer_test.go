package lexer

import (
	"testing"

	"github.com/Selyss/AutoCal/token"
)

func TestNextToken(t *testing.T) {
	input := `
	IF DAY EVEN, ADD "quite an odd day!".

	IF DAY ODD, ADD (
		TITLE "Hitting the gym!",
		DESC "Chest!!!",
		RED
	).
	  
	ELSE, ADD (
		TITLE "Rest day.",
		GREY
	).
	
	LET commute = (
		TITLE "Bus",
		DESC "take the normal bus route",
		TIME,
		RED
	).
	  
	IF DAY IS 1, ADD commute("9am").`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "IF"},
		{token.DAY, "DAY"},
		{token.EVEN, "EVEN"},
		{token.COMMA, ","},
		{token.ADD, "ADD"},
		{token.STRING, "quite an odd day!"},
		{token.PERIOD, "."},
		{token.IF, "IF"},
		{token.DAY, "DAY"},
		{token.ODD, "ODD"},
		{token.COMMA, ","},
		{token.ADD, "ADD"},
		{token.LPAREN, "("},
		{token.TITLE, "TITLE"},
		{token.STRING, "Hitting the gym!"},
		{token.COMMA, ","},
		{token.DESC, "DESC"},
		{token.STRING, "Chest!!!"},
		{token.COMMA, ","},

		{token.RED, "RED"},
		{token.RPAREN, ")"},
		{token.PERIOD, "."},
		{token.ELSE, "ELSE"},
		{token.COMMA, ","},
		{token.ADD, "ADD"},
		{token.LPAREN, "("},
		{token.TITLE, "TITLE"},
		{token.STRING, "Rest day."},
		{token.COMMA, ","},
		{token.GREY, "GREY"},
		{token.RPAREN, ")"},
		{token.PERIOD, "."},
		{token.LET, "LET"},
		{token.IDENT, "commute"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.TITLE, "TITLE"},
		{token.STRING, "Bus"},
		{token.COMMA, ","},
		{token.DESC, "DESC"},
		{token.STRING, "take the normal bus route"},
		{token.COMMA, ","},
		{token.TIME, "TIME"},
		{token.COMMA, ","},
		{token.RED, "RED"},
		{token.RPAREN, ")"},
		{token.PERIOD, "."},
		{token.IF, "IF"},
		{token.DAY, "DAY"},
		{token.EQ, "IS"},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.ADD, "ADD"},
		{token.IDENT, "commute"},
		{token.LPAREN, "("},
		{token.STRING, "9am"},
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
