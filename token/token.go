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
	LET    = "LET"
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Colors
	RED  = "RED"
	GREY = "GREY"

	// Operators
	ASSIGN = "="
	EQ     = "IS"
	NOT_EQ = "IS NOT"

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
	"RED":  RED,
	"GREY": GREY,

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
	"IS":     EQ,
	"IS NOT": NOT_EQ,
	"LET":    LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
