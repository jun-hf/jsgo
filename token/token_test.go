package token

import (
	"testing"
)

func TestTokenString(t *testing.T) {
	tests := []struct {
		token    Token
		expected string
	}{
		{ILLEGAL, "ILLEGAL"},
		{EOF, "EOF"},
		{IDENT, "IDENTIFIER"},
		{NUMBER, "NUMBER"},
		{STRING, "STRING"},
		{ADD, "+"},
		{MINUS, "-"},
		{MUL, "*"},
		{DIVIDE, "/"},
		{LSS, "<"},
		{GTR, ">"},
		{BANG, "!"},
		{ASSIGN, "="},
		{NOT_EQUAL, "!="},
		{EQUAL, "=="},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{DOT, "."},
		{COLON, ":"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{FUNCTION, "function"},
		{VAR, "var"},
		{IF, "if"},
		{ELSE, "else"},
		{ELSEIF, "elseif"},
		{RETURN, "return"},
		{TRUE, "true"},
		{FALSE, "false"},
	}

	for _, tt := range tests {
		if tt.token.String() != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, tt.token.String())
		}
	}
}

func TestIsLiteral(t *testing.T) {
	if !IDENT.IsLiteral() || !NUMBER.IsLiteral() || !STRING.IsLiteral() {
		t.Errorf("expected true for literals")
	}

	if ADD.IsLiteral() {
		t.Errorf("expected false for non-literals")
	}
}

func TestIsOperator(t *testing.T) {
	if !ADD.IsOperator() || !MINUS.IsOperator() || !MUL.IsOperator() {
		t.Errorf("expected true for operators")
	}

	if IDENT.IsOperator() {
		t.Errorf("expected false for non-operators")
	}
}

func TestIsKeyword(t *testing.T) {
	keywords := []string{"function", "var", "if", "else", "elseif", "return", "true", "false"}

	for _, keyword := range keywords {
		if !IsKeyword(keyword) {
			t.Errorf("expected true for keyword %s", keyword)
		}
	}

	nonKeywords := []string{"apple", "banana", "IDENT", "NUMBER"}

	for _, nonKeyword := range nonKeywords {
		if IsKeyword(nonKeyword) {
			t.Errorf("expected false for non-keyword %s", nonKeyword)
		}
	}
}

func TestIsIdentifier(t *testing.T) {
	identifiers := []string{"apple", "banana", "myVar"}

	for _, identifier := range identifiers {
		if !IsIdentifier(identifier) {
			t.Errorf("expected true for identifier %s", identifier)
		}
	}

	nonIdentifiers := []string{"", "var", "if"}

	for _, nonIdentifier := range nonIdentifiers {
		if IsIdentifier(nonIdentifier) {
			t.Errorf("expected false for non-identifier %s", nonIdentifier)
		}
	}
}
