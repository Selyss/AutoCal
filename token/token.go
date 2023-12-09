package token

type TokenType string

const (
	IF      TokenType = "IF"
	DAY     TokenType = "DAY"
	ODD     TokenType = "ODD"
	EVEN    TokenType = "EVEN"
	ADD     TokenType = "ADD"
	TIME    TokenType = "TIME"
	COMMA   TokenType = ","
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"if": IF,
}
