package parser

import (
	"fmt"
	"orca/ast"
	"orca/lexer"
	"orca/token"
)

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(token token.TokenType) {
	// TODO: support line numbers and full report on code
	msg := fmt.Sprintf("expected next token to be %s, got %s instead.", token, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) expectPeek(token token.TokenType) bool {
	if p.peekToken.Type == token {
		p.nextToken()
		return true
	}
	p.peekError(token)
	return false
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// Skip until simicolon temporarily
	// TODO: update this when expression parser is implemented
	for !(p.currToken.Type == token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}
