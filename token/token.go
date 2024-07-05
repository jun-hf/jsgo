package token

import (
	"strconv"
)

// Token is the set of tokens in the jsgo language
type Token int

const (
	ILLEGAL Token = iota
	EOF 
	
	literalBegin

	IDENT // apple
	NUMBER // 89
	STRING // "number"

	literalEnd

	operatorBegin // operator and delimiter

	ADD // +
	MINUS // -
	MUL // *
	DIVIDE // /

	LSS // < 
	GTR // >
	BANG // !
	ASSIGN // =

	NOT_EQUAL // !=
	EQUAL // ==

	COMMA // ,
	SEMICOLON // ;
	DOT // .
	COLON // :

	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

	operatorEnd

	keywordBegin // keyword in the language of jsgo

	FUNCTION // function
	VAR // var
	IF // if
	ELSE // else
	ELSEIF // elseif
	RETURN // return
	TRUE // true
	FALSE // false

	keywordEnd

)

var keywords = map[string]Token{
	"var": VAR,
	"function": FUNCTION,
	"if": IF,
	"else": ELSE,
	"elseif": ELSEIF,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

// tokens store the repective string representation of the token
var tokens = [...]string{
	ILLEGAL:    "ILLEGAL",
	EOF:        "EOF",
	IDENT:      "IDENTIFIER",
	NUMBER:     "NUMBER",
	STRING:     "STRING",
	ADD:        "+",
	MINUS:      "-",
	MUL:        "*",
	DIVIDE:     "/",
	LSS:        "<",
	GTR:        ">",
	BANG:       "!",
	ASSIGN:     "=",
	NOT_EQUAL:  "!=",
	EQUAL:      "==",
	COMMA:      ",",
	SEMICOLON:  ";",
	DOT:        ".",
	COLON:      ":",
	LPAREN:     "(",
	RPAREN:     ")",
	LBRACE:     "{",
	RBRACE:     "}",
	FUNCTION:   "function",
	VAR:        "var",
	IF:         "if",
	ELSE:       "else",
	ELSEIF:     "elseif",
	RETURN:     "return",
	TRUE:       "true",
	FALSE:      "false",
}

// String return the string representation of the token
func (t Token) String() string {
	s := ""
	if t >= 0 &&  int(t) < len(tokens) {
		s = tokens[t]
	}

	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ") does not exist"
	}
	return s
}

// IsLiteral determine if the token is a literal in the language 
func (t Token) IsLiteral() bool { return literalBegin < t && t < literalEnd }

// IsOperator determine if the token is an operator in the language
func (t Token) IsOperator() bool { return operatorEnd < t && t < operatorEnd }

// IsKeyword determine if the token is a keyword
func (t Token) IsKeyword() bool {
	s := tokens[t]
	return IsKeyword(s)
}

// IsKeyword determine if the string pass into it is part of a keyword in the language
func IsKeyword(word string) bool {
	_, ok := keywords[word]
	return ok
}

// IsIdentifier determine if the string pass into it is an identifier
func IsIdentifier(word string) bool {
	if len(word) == 0 || IsKeyword(word) {
		return false
	}
	return true
}