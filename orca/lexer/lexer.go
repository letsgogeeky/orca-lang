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

func (l *Lexer) getOperator() token.Token {
	var t token.Token
	ch := l.currChar
	switch ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			literal := string(ch) + string(l.currChar)
			return token.Token{Type: token.EQ, Literal: literal}
		}
		t = newToken(token.ASSIGN, ch)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			literal := string(ch) + string(l.currChar)
			return token.Token{Type: token.NOT_EQ, Literal: literal}
		}
		t = newToken(token.BANG, ch)
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			literal := string(ch) + string(l.currChar)
			return token.Token{Type: token.GT_EQ, Literal: literal}
		}
		t = newToken(token.GT, ch)
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			literal := string(ch) + string(l.currChar)
			return token.Token{Type: token.LT_EQ, Literal: literal}
		}
		t = newToken(token.LT, ch)
	}
	return t
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.consumeWhiteSpace()
	switch l.currChar {
	case '=', '!', '<', '>':
		t = l.getOperator()
	case '+':
		t = newToken(token.PLUS, l.currChar)
	case '-':
		t = newToken(token.MINUS, l.currChar)
	case '/':
		t = newToken(token.SLASH, l.currChar)
	case '*':
		t = newToken(token.ASTERISK, l.currChar)
	case '^':
		t = newToken(token.POW, l.currChar)
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
	default:
		if isLetter(l.currChar) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		}
		if isDigit(l.currChar) {
			t.Literal = l.readInt()
			t.Type = token.INT
			return t
		}
		t = newToken(token.ILLEGAL, l.currChar)
	}
	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	// TODO: support multitype literals
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) consumeWhiteSpace() {
	for l.currChar == ' ' || l.currChar == '\t' || l.currChar == '\n' || l.currChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readInt() string {
	position := l.position
	for isDigit(l.currChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= l.inputLen {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	// Extend this function to define and support different ways
	// for identifiers and values
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
