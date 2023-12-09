package main

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

func NewToken(Type TokenType, Literal string) Token {
	return Token{Type: Type, Literal: Literal}
}
