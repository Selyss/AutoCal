package token

type TokenType string

const (
	COMMA   = ","
	PERIOD  = "."
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"

	// Keywords
	IF     = "IF"
	ELSE   = "ELSE"
	DAY    = "DAY"
	ODD    = "ODD"
	EVEN   = "EVEN"
	ADD    = "ADD"
	STRING = "STRING"
	TIME   = "TIME"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"IF":     IF,
	"ELSE":   ELSE,
	"DAY":    DAY,
	"ODD":    ODD,
	"EVEN":   EVEN,
	"ADD":    ADD,
	"STRING": STRING,
	"TIME":   TIME,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
