package tokenizer

import (
	"fmt"
)

type TokenType string

const (
	TOKEN_PRINT       TokenType = "PRINT"
	TOKEN_LET         TokenType = "LET"
	TOKEN_IF          TokenType = "IF"
	TOKEN_THEN        TokenType = "THEN"
	TOKEN_ELSE        TokenType = "ELSE"
	TOKEN_WHILE       TokenType = "WHILE"
	TOKEN_DO          TokenType = "DO"
	TOKEN_STOP        TokenType = "STOP"
	TOKEN_END         TokenType = "END"
	TOKEN_IDENTIFIER  TokenType = "IDENTIFIER"
	TOKEN_INTEGER     TokenType = "INTEGER"
	TOKEN_FLOAT       TokenType = "FLOAT"
	TOKEN_COMMENT     TokenType = "COMMENT"
	TOKEN_ADD_SUB     TokenType = "ADD_SUB_OPERATOR"
	TOKEN_MUL_DIV     TokenType = "MUL_DIV_OPERATOR"
	TOKEN_REL_OP      TokenType = "RELATIONAL_OPERATOR"
	TOKEN_EQUALS      TokenType = "EQUALS_OPERATOR"
	TOKEN_LEFT_PAREN  TokenType = "LEFT_PAREN"
	TOKEN_RIGHT_PAREN TokenType = "RIGHT_PAREN"
	TOKEN_EOF         TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

var keywords = map[string]TokenType{
	"PRINT": TOKEN_PRINT,
	"LET":   TOKEN_LET,
	"IF":    TOKEN_IF,
	"THEN":  TOKEN_THEN,
	"ELSE":  TOKEN_ELSE,
	"WHILE": TOKEN_WHILE,
	"DO":    TOKEN_DO,
	"STOP":  TOKEN_STOP,
	"END":   TOKEN_END,
}

var operators = map[string]TokenType{
	"+":  TOKEN_ADD_SUB,
	"-":  TOKEN_ADD_SUB,
	"*":  TOKEN_MUL_DIV,
	"/":  TOKEN_MUL_DIV,
	"=":  TOKEN_EQUALS,
	"==": TOKEN_REL_OP,
	"<":  TOKEN_REL_OP,
	">":  TOKEN_REL_OP,
}

type TokenizerError struct {
	Position int
	Char     rune
	Message  string
}

func (e *TokenizerError) Error() string {
	return fmt.Sprintf("Tokenizer Error at position %d: Unexpected character '%c'. %s", e.Position, e.Char, e.Message)
}
