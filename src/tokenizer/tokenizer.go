package tokenizer

import (
	"unicode"
)

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	i := 0
	runes := []rune(input)

	for i < len(runes) {
		ch := runes[i]

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		// Handle numbers
		if unicode.IsDigit(ch) {
			start := i
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				i++
			}

			if i < len(runes) && runes[i] == '.' {
				i++
				for i < len(runes) && unicode.IsDigit(runes[i]) {
					i++
				}
				tokens = append(tokens, Token{Type: TOKEN_FLOAT, Value: string(runes[start:i])})
				continue
			}

			tokens = append(tokens, Token{Type: TOKEN_INTEGER, Value: string(runes[start:i])})
			continue
		}

		// Handle comments
		const remKeyword = "REM"
		if i+len(remKeyword) <= len(runes) && string(runes[i:i+3]) == remKeyword {
			i += 3
			start := i

			for i < len(runes) && runes[i] != '\n' {
				i++
			}
			tokens = append(tokens, Token{Type: TOKEN_COMMENT, Value: string(runes[start : i-1])})
			continue
		}

		// Handle keywords and identifiers
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

		// Handle strings
		if ch == '"' {
			start := i
			i++

			for i < len(runes) && runes[i] != '"' {
				i++
			}

			if i < len(runes) && runes[i] == '"' {
				i++
				tokens = append(tokens, Token{Type: TOKEN_STRING, Value: string(runes[start+1 : i-1])})
			} else {
				return nil, &TokenizerError{Position: i, Char: ch, Message: "Unclosed string literal."}
			}
			continue
		}

		// Handle operators
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
