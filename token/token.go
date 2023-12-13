package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Colors
	RED   = "RED"
	BLUE  = "BLUE"
	GREEN = "GREEN"
	GREY  = "GREY"

	// Operators
	GThan  = ">"
	LThan  = "<"
	ASSIGN = "="
	IS     = "IS" // assign

	// Delimiters
	COMMA  = ","
	PERIOD = "."
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	IF    = "IF"
	ELSE  = "ELSE"
	DAY   = "DAY"
	ODD   = "ODD"
	EVEN  = "EVEN"
	ADD   = "ADD"
	TIME  = "TIME"
	TITLE = "TITLE"
	DESC  = "DESC"
)

var keywords = map[string]TokenType{
	"RED":   RED,
	"BLUE":  BLUE,
	"GREEN": GREEN,
	"GREY":  GREY,

	"IF":     IF,
	"ELSE":   ELSE,
	"DAY":    DAY,
	"ODD":    ODD,
	"EVEN":   EVEN,
	"ADD":    ADD,
	"STRING": STRING,
	"TIME":   TIME,
	"TITLE":  TITLE,
	"DESC":   DESC,
	"IS":     IS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
