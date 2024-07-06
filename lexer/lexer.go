package lexer

import (
	"strconv"
	"strings"

	"github.com/jun-hf/jsgo/token"
)

// In the lexer we should have an record of the error recording the position of the file
// Parser should have an record of the error that indicate the user of the error
type Lexer struct {
	src []byte
	position int
	nextPosition int
	ch byte
	errors []string
}

func New(byt []byte) *Lexer {
	l := &Lexer{ src: byt }
	l.next()
	return l
}

// In this lexer consider how to report error to the users
func (l *Lexer) Lex() (token.Token, string) {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '+':
		tok = token.ADD
	case '-':
		tok = token.MINUS
	case '*':
		tok = token.MUL
	case '/':
		tok = token.DIVIDE
	case ',':
		tok = token.COMMA
	case '.':
		tok = token.DOT
	case ':':
		tok = token.COLON
	case ';':
		tok = token.SEMICOLON
	case '(':
		tok = token.LPAREN
	case ')':
		tok = token.RPAREN
	case '{':
		tok = token.LBRACE
	case '}':
		tok = token.RBRACE
	case '!':
		if l.peekByte() == '=' {
			l.next()
			tok = token.NOT_EQUAL
			break
		}
		tok = token.BANG
	case '=':
		if l.peekByte() == '=' {
			l.next()
			tok = token.EQUAL
			break
		}
		tok = token.ASSIGN
	case 0:
		return token.EOF, token.EOF.String()
	default:
		if l.isLetter() {
			s := l.getLetter()
			if token.IsKeyword(s) {
				keyword, _ := token.Keyword(s) // ignoring error because the previous if checks for keyword
				return keyword, s
			}
			return token.IDENT, s
		}
		if l.isDigit() {
			return l.getDigitToken()
		}
		return token.ILLEGAL, token.ILLEGAL.String()
	}
	l.next()
	return tok, tok.String()
}

func (l *Lexer) getLetter() string {
	var letter strings.Builder
	for l.isLetter() {
		letter.WriteByte(l.ch)
		l.next()
	}
	return letter.String()
}

// getDigitToken returns either [token.Token.NUMBER] or [token.Token.FLOAT] 
// with its corresponding literal
func (l *Lexer) getDigitToken() (token.Token, string) {
	var digit strings.Builder
	hasDot := false
	for {
		if l.isDigit() {
			digit.WriteByte(l.ch)
			l.next()
		} else if l.ch == '.' {
			if hasDot {
				l.errors = append(l.errors, "digit formatted incorrect at " + strconv.Itoa(l.position))
			}
			hasDot = true
			digit.WriteByte(l.ch)
			l.next()
		} else {
			break
		}
	}
	if hasDot {
		return token.FLOAT, digit.String()
	}
	return token.NUMBER, digit.String()
}

func (l *Lexer) isLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || l.ch == '_' ||'A' <= l.ch && l.ch <= 'Z' 
}

func (l *Lexer) isDigit() bool {
	return '0' <= l.ch && l.ch <= '9'
}

// next moves the current position of the char in [Lexer.data] to the next one
// it will the [Lexer.ch] to 0 when [Lexer.position] is at the last byte of the [Lexer.src]
func (l *Lexer) next() {
	if l.nextPosition == len(l.src) {
		l.ch = 0
		l.position = l.nextPosition
		return 
	} 
	l.position = l.nextPosition
	l.ch = l.src[l.position]
	l.nextPosition++
}

// peekByte returns the next byte in [Lexer.src] without moviing the [Lexer.position] and [Lexer.NextPostion]
// use this function to look ahead of the [Lexer.src]
func (l *Lexer) peekByte() byte {
	if l.nextPosition == len(l.src) {
		return 0 // 0 represent End of file
	}
	return l.src[l.nextPosition] 
}

// skipWhitespace skips all the current whitespace in [Lexer.src]
func (l *Lexer) skipWhitespace() {
	for  l.ch == ' ' || l.ch  == '\t' || l.ch  == '\n' || l.ch  == '\r' {
		l.next()
	}
}