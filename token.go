package elmlexer

type Token struct {
	Type    TokenType
	Pos     int
	Value   string
	Line    int
	LinePos int
}

// NewToken creates a new item of type t with start, line, value and position in line.
func NewToken(t TokenType, start, linePos, line int, val string) *Token {
	return &Token{
		Type:    t,
		Pos:     start,
		Value:   val,
		LinePos: linePos,
		Line:    line,
	}
}

//go:generate stringer -type=TokenType

// TokenType is the type of an item.
type TokenType uint

const (
	// TokenError is an error occurred in the process of lexing, value is the text of the error
	TokenError TokenType = iota
	// TokenEOF is the end of the input
	TokenEOF
	// TokenComment is an user comment
	TokenComment
	// TokenLeftParen is the left parenthesis "("
	TokenLeftParen
	// TokenRightParen is the right parenthesis ")"
	TokenRightParen
	// TokenLeftBracket is the left bracket "["
	TokenLeftBracket
	// TokenRightBracket is the right bracket "]"
	TokenRightBracket
	// TokenLeftBrace is the left brace "{"
	TokenLeftBrace
	// TokenRightBrace is the right brace "}"
	TokenRightBrace
	// TokenPipe is the pipe character "|"
	TokenPipe
	// TokenInfixOp is an identifier between backticks that acts as an infix op
	TokenInfixOp
	// TokenColon is the colon character ":"
	TokenColon
	// TokenAssign is the equal character "="
	TokenAssign
	// TokenComma is the comma character ","
	TokenComma
	// TokenArrow is the arrow operator "->"
	TokenArrow
	// TokenIdentifier is an identifier (user defined vars, functions, predefined, ...)
	TokenIdentifier
	// TokenOp is an operator
	TokenOp
	// TokenString is a quoted string literal
	TokenString
	// TokenInt is an integer number
	TokenInt
	// TokenFloat is a floating point number
	TokenFloat
	// TokenKeyword is a keyword ":keyword"
	TokenKeyword
	// TokenRange is a range of integers
	TokenRange
	// TokenChar is a quoted character literal
	TokenChar
	// TokenBool is a boolean (true or false)
	TokenBool
	// TokenDot is the dot character "."
	TokenDot
)
