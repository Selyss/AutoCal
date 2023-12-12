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
	LPAREN  = "("
	RPAREN  = ")"
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
	TITLE  = "TITLE"
	DESC   = "DESC"

	// Colors
	RED   = "RED"
	BLUE  = "BLUE"
	GREEN = "GREEN"
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
	"TITLE":  TITLE,
	"DESC":   DESC,
	"RED":    RED,
	"BLUE":   BLUE,
	"GREEN":  GREEN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
