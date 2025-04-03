package tokenizer

import (
	"unicode"
)

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	i := 0
	runes := []rune(input)
	line := 1

	for i < len(runes) {
		ch := runes[i]

		if ch == '\n' {
			line++
		}

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		// Handle numbers
		if unicode.IsDigit(ch) || ch == '-' {
			start := i

			if ch == '-' {
				i++
			}

			for i < len(runes) && unicode.IsDigit(runes[i]) {
				i++
			}

			if i < len(runes) && runes[i] == '.' {
				i++
				for i < len(runes) && unicode.IsDigit(runes[i]) {
					i++
				}
				tokens = append(tokens, Token{Type: TOKEN_FLOAT, Value: string(runes[start:i]), Line: line})
				continue
			}

			tokens = append(tokens, Token{Type: TOKEN_INTEGER, Value: string(runes[start:i]), Line: line})
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
			tokens = append(tokens, Token{Type: TOKEN_COMMENT, Value: string(runes[start : i-1]), Line: line})
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
				tokens = append(tokens, Token{Type: tokenType, Value: word, Line: line})
			} else {
				tokens = append(tokens, Token{Type: TOKEN_IDENTIFIER, Value: word, Line: line})
			}
			continue
		}

		// Handle ==
		if ch == '=' && i+1 < len(runes) && runes[i+1] == '=' {
			tokens = append(tokens, Token{Type: TOKEN_REL_OP, Value: "==", Line: line})
			i += 2
			continue
		}

		// Handle operators
		op := string(ch)
		if tokenType, found := operators[op]; found {
			tokens = append(tokens, Token{Type: tokenType, Value: op, Line: line})
			i++
			continue
		}

		// Handle parentheses
		if ch == '(' {
			tokens = append(tokens, Token{Type: TOKEN_LEFT_PAREN, Value: string(ch), Line: line})
			i++
			continue
		}
		if ch == ')' {
			tokens = append(tokens, Token{Type: TOKEN_RIGHT_PAREN, Value: string(ch), Line: line})
			i++
			continue
		}

		return nil, &TokenizerError{Position: i, Char: ch, Message: "Unknown token encountered."}

	}

	tokens = append(tokens, Token{Type: TOKEN_EOF, Value: "EOF", Line: line})
	return tokens, nil
}
