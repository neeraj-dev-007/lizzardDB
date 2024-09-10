package parser

import "strings"

type TokenType int

// Go’s iota identifier is used in const declarations to simplify definitions of incrementing numbers.
// The value of iota is reset to 0 whenever the reserved word const appears in the source (i.e. each const block)
// incremented by one after each ConstSpec e.g. each Line.
const (
	tokenEOF TokenType = iota
	tokenKeyword
	tokenIdentifier
	tokenOperator
	tokenNumber
	tokenString
	tokenComma
	tokenSemicolon
	tokenParenthesis
	tokenError
)

// Token represents a single token
type Token struct {
	tokenType TokenType
	value     string
}

type Lexer struct {
	input string
	start int
	pos   int
	width int
}

func Tokenize(sql string) {
	tokens := []Token{}
	words := strings.Fields(sql)
	for _, word := range words {
		switch word {
		case "SELECT": 
			tokens = append(tokens, Token{tokenType: tokenKeyword, value: word})
		case "FROM":
			tokens = append(tokens, Token{tokenType: tokenKeyword, value: word})
		}
	}
}