package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // variable_name, function_name
	INT   = "INT"
	FLOAT = "FLOAT"
	HEX   = "HEX"
	OCT   = "OCT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	POW      = "^"
	BANG     = "!"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	LT_EQ = "<="
	GT_EQ = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	L_PAREN = "("
	R_PAREN = ")"

	L_BRACE = "{"
	R_BRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	BREACH   = "BREACH" // try .. (Exception handling)
	CATCH    = "CATCH"
)

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"let":     LET,
	"if":      IF,
	"else":    ELSE,
	"true":    TRUE,
	"false":   FALSE,
	"return":  RETURN,
	"whistle": FUNCTION,
	"breach":  BREACH,
	"catch":   CATCH,
}

func LookupIdent(iden string) TokenType {
	if tok, ok := keywords[iden]; ok {
		return tok
	}
	return IDENT
}
