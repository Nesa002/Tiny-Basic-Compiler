package tokenizer

import (
	"fmt"
)

type TokenType string

const (
	TOKEN_PRINT      TokenType = "PRINT"
	TOKEN_LET        TokenType = "LET"
	TOKEN_IF         TokenType = "IF"
	TOKEN_THEN       TokenType = "THEN"
	TOKEN_END        TokenType = "END"
	TOKEN_IDENTIFIER TokenType = "IDENTIFIER"
	TOKEN_INTEGER    TokenType = "INTEGER"
	TOKEN_FLOAT      TokenType = "FLOAT"
	TOKEN_STRING     TokenType = "STRING"
	TOKEN_COMMENT    TokenType = "COMMENT"
	TOKEN_OPERATOR   TokenType = "OPERATOR"
	TOKEN_REL_OP     TokenType = "RELATIONAL_OPERATOR"
	TOKEN_EOF        TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

var keywords = map[string]TokenType{
	"PRINT": TOKEN_PRINT,
	"LET":   TOKEN_LET,
	"IF":    TOKEN_IF,
	"THEN":  TOKEN_THEN,
	"END":   TOKEN_END,
}

var operators = map[string]TokenType{
	"+": TOKEN_OPERATOR,
	"-": TOKEN_OPERATOR,
	"*": TOKEN_OPERATOR,
	"/": TOKEN_OPERATOR,
	"=": TOKEN_REL_OP,
	"<": TOKEN_REL_OP,
	">": TOKEN_REL_OP,
}

type TokenizerError struct {
	Position int
	Char     rune
	Message  string
}

func (e *TokenizerError) Error() string {
	return fmt.Sprintf("Tokenizer Error at position %d: Unexpected character '%c'. %s", e.Position, e.Char, e.Message)
}
