package token

type TokenType string

const (
	COMMA   TokenType = ","
	PERIOD  TokenType = "."
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Keywords
	IF     TokenType = "IF"
	DAY    TokenType = "DAY"
	ODD    TokenType = "ODD"
	EVEN   TokenType = "EVEN"
	ADD    TokenType = "ADD"
	STRING TokenType = "STRING"
	TIME   TokenType = "TIME"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"IF":     IF,
	"DAY":    DAY,
	"ODD":    ODD,
	"EVEN":   EVEN,
	"ADD":    ADD,
	"STRING": STRING,
	"TIME":   TIME,
}
