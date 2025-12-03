package token

type TokenType string

type Token struct {
	Type     TokenType
	Literral string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN         = "="
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"
	//key words
	PRINT = "PRINT"
	INPUT = "INPUT"
	IF    = "IF"
	JUMP  = "JUMP"
	LABLE = "LABLE"
)
