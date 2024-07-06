package lexer

import (
	"testing"

	"github.com/jun-hf/jsgo/token"
)

func TestLex(t *testing.T) {
	tests := []struct{
		input string
		expected token.Token
		expectedLit string
	}{
		{"+", token.ADD, "+"},
		{"-", token.MINUS, "-"},
		{"*", token.MUL, "*"},
		{"/", token.DIVIDE, "/"},
		{",", token.COMMA, ","},
		{".", token.DOT, "."},
		{":", token.COLON, ":"},
		{";", token.SEMICOLON, ";"},
		{"(", token.LPAREN, "("},
		{")", token.RPAREN, ")"},
		{"{", token.LBRACE, "{"},
		{"}", token.RBRACE, "}"},
		{"=", token.ASSIGN, "="},
		{"!", token.BANG, "!"},
		{ "!=", token.NOT_EQUAL, "!="},
		{"==", token.EQUAL, "=="},
		{"", token.EOF, "EOF"},
		{"89", token.NUMBER, "89"},
		{"hello", token.IDENT, "hello"},
		{"89.2", token.FLOAT, "89.2"},
		{"var", token.VAR, "var"},
		{"function", token.FUNCTION, "function"},
		{"if", token.IF, "if"},
		{"else", token.ELSE, "else"},
		{"elseif", token.ELSEIF, "elseif"},
		{"return", token.RETURN, "return"},
		{"false", token.FALSE, "false"},
		{"true", token.TRUE, "true"},
	}

	for _, test := range tests {
		l := New([]byte(test.input))
		tok, lit := l.Lex()
		if tok != test.expected {
			t.Errorf("Lexer.Lex wrong token. got=%q, expected=%q", tok, test.expected)
		}
		if lit != test.expectedLit {
			t.Errorf("Lexer.Lex wrong lit. got=%q, expected=%q", lit, test.expectedLit)
		}
	}
}

// func TestScan(t *testing.T) {
// 	data, err := os.ReadFile("./example.gs")
// 	if err != nil {
// 		t.Fatal("faild to open: example.gs")
// 	}
// 	l := New(data)
// 	l.Next()
// }

// const (
// 	special = iota
// 	literal 
// 	operator
// 	keyword
// )

// var tokens = []struct{
// 	tok token.Token
// 	lit string
// 	class int
// }{
// 	{token.IDENT, "bar", literal},
// 	{token.IDENT, "apple", literal},
// 	{token.NUMBER, "89", literal},
// 	{token.ADD, "+", operator},
// 	{token.MINUS, "-", operator},
// 	{token.MUL, "*", operator},
// 	{token.DIVIDE, "/", operator},
// }

// const whitespace = "\n" // separate token
// var source = func() []byte {
// 	var src []byte
// 	for _, t := range tokens {
// 		src = append(src, t.lit...)
// 		src = append(src, whitespace...)
// 	}
// 	return src
// }