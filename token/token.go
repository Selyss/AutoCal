package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	COMMA   = ","
	PERIOD  = "."
	COLON   = ":"
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
	COLOR  = "COLOR"
)

var keywords = map[string]TokenType{
	"IF":     IF,
	"ELSE":   ELSE,
	"DAY":    DAY,
	"ODD":    ODD,
	"EVEN":   EVEN,
	"ADD":    ADD,
	"STRING": STRING,
	"TIME":   TIME,
	"COLOR":  COLOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
