package parser

import (
	"orca/ast"
	"orca/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let gawava = 23234234;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	hasErrors := hasParserErrors(t, p)
	if hasErrors {
		panic("Parser has unexpected errors")
	}
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements: Expected 3 statements. Got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"gawava"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestInvalidLetStatements(t *testing.T) {
	input := `
let x 5;
let y = 10;
let gawava = 23234234;
`

	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()
	hasErrors := hasParserErrors(t, p)
	if !hasErrors {
		panic("Expected parser errors but got no errors!")
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral(): Expected value 'let', Got: %q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not of type '*ast.LetStatement'. got: %T", statement)
		return false
	}
	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value expected '%s' got: '%s'", name, letStatement.Name.Value)
		return false
	}
	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.TokenLiteral() expected '%s'. got '%s'", name, letStatement.Name.TokenLiteral())
	}
	return true
}

func hasParserErrors(t *testing.T, p *Parser) bool {
	errors := p.Errors()

	if len(errors) == 0 {
		return false
	}

	t.Logf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Logf("Parser error: %q", msg)
	}
	return true
}
