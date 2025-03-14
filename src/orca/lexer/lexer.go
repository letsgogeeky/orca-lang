package lexer

import "orca/token"

// TODOS:
// 1- Support Unicode input and emojis
// 2- Support multitype literals

type Lexer struct {
	input        string
	position     int // a responsible pointer to current character
	readPosition int // a rabbit peeking into next character
	currChar     byte
	inputLen     int // a precomputed input length to avoid calling len in iteration
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.inputLen = len(input)
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= l.inputLen {
		l.currChar = 0
	} else {
		l.currChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.currChar {
	case '=':
		t = newToken(token.ASSIGN, l.currChar)
	case '+':
		t = newToken(token.PLUS, l.currChar)
	case '-':
		t = newToken(token.MINUS, l.currChar)
	case '/':
		t = newToken(token.SLASH, l.currChar)
	case '*':
		t = newToken(token.ASTERISK, l.currChar)
	case '(':
		t = newToken(token.L_PAREN, l.currChar)
	case ')':
		t = newToken(token.R_PAREN, l.currChar)
	case '{':
		t = newToken(token.L_BRACE, l.currChar)
	case '}':
		t = newToken(token.R_BRACE, l.currChar)
	case ',':
		t = newToken(token.COMMA, l.currChar)
	case ';':
		t = newToken(token.SEMICOLON, l.currChar)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	}
	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	// TODO: support multitype literals
	return token.Token{Type: tokenType, Literal: string(ch)}
}
