package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"

	// Identifiers + literals
	IDENT    = "IDENT" // variable_name, function_name
	INT	     = "INT"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	BANG      = "!"

	LT        = "<"
	GT        = ">"

	EQ        = "=="
	NOT_EQ    = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	L_PAREN   = "("
	R_PAREN   = ")"

	L_BRACE   = "{"
	R_BRACE   = "}"

	// keywords
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"

)