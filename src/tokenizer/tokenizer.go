package tokenizer

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType string

const (
	TOKEN_PRINT      TokenType = "PRINT"
	TOKEN_LET        TokenType = "LET"
	TOKEN_IF         TokenType = "IF"
	TOKEN_THEN       TokenType = "THEN"
	TOKEN_REM        TokenType = "REM"
	TOKEN_END        TokenType = "END"
	TOKEN_IDENTIFIER TokenType = "IDENTIFIER"
	TOKEN_NUMBER     TokenType = "NUMBER"
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
	"REM":   TOKEN_REM,
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

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	i := 0
	runes := []rune(strings.ToUpper(input))

	for i < len(runes) {
		ch := runes[i]

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		if unicode.IsDigit(ch) {
			start := i
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				i++
			}

			tokens = append(tokens, Token{Type: TOKEN_NUMBER, Value: string(runes[start:i])})
			continue
		}

		if unicode.IsLetter(ch) {
			start := i
			for i < len(runes) && unicode.IsLetter(runes[i]) {
				i++
			}

			word := string(runes[start:i])

			if tokenType, found := keywords[word]; found {
				tokens = append(tokens, Token{Type: tokenType, Value: word})
			} else {
				tokens = append(tokens, Token{Type: TOKEN_IDENTIFIER, Value: word})
			}
			continue
		}

		op := string(ch)

		if tokenType, found := operators[op]; found {
			tokens = append(tokens, Token{Type: tokenType, Value: op})
			i++
			continue
		}

		return nil, &TokenizerError{Position: i, Char: ch, Message: "Unknown token encountered."}

	}

	tokens = append(tokens, Token{Type: TOKEN_EOF, Value: "EOF"})
	return tokens, nil
}
