package lexer

import (
	"testing"

	"orca/token"
)

type expectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func runComparisons(input string, tests []expectedToken, t *testing.T) {
	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - incorrect token type. expected=%q, got=%q",
				i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - incorrect token literal. expected=%q, got=%q",
				i, tt.expectedLiteral, token.Literal)
		}
	}
}
func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []expectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.L_PAREN, "("},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.R_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runComparisons(input, tests, t)
}

func TestNextTokenExtended(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
}

let result = add(five, ten);
`
	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.L_PAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.L_PAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.R_PAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	runComparisons(input, tests, t)
}

func TestOperatorTokens(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
}

let result = add(five, ten);
!-/*^5;
5 < 10 > 5;
`
	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.L_PAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.L_PAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.R_PAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.POW, "^"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	runComparisons(input, tests, t)
}

func TestKeywordTokens(t *testing.T) {
	input := `let five = 5;
if (5 < 10) {
	return true;
} else {
	return false;
}
`
	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.L_PAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.ELSE, "else"},
		{token.L_BRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.EOF, ""},
	}
	runComparisons(input, tests, t)
}

func TestFunctionTokens(t *testing.T) {
	input := `let pow = whistle(x, y) {
	return x ^ y;
}
`
	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENT, "pow"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "whistle"},
		{token.L_PAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.POW, "^"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.EOF, ""},
	}

	runComparisons(input, tests, t)
}

func TestComparisonOperatorTokens(t *testing.T) {
	input := `let age = 12;
age == 14;
age != 0;
age >= 0;
age <= 0;
age > 0;
age < 0;
`

	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENT, "age"},
		{token.ASSIGN, "="},
		{token.INT, "12"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "age"},
		{token.EQ, "=="},
		{token.INT, "14"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "age"},
		{token.NOT_EQ, "!="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "age"},
		{token.GT_EQ, ">="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "age"},
		{token.LT_EQ, "<="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "age"},
		{token.GT, ">"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "age"},
		{token.LT, "<"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runComparisons(input, tests, t)
}
