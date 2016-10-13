package elmlexer

import (
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAllowedInIdentifier(t *testing.T) {
	allowed := "abc135fdcv_'"
	notAllowed := ":.;,{}[]`|#%\\-+-?!&=/<>^$"
	for _, r := range allowed {
		assert.Equal(t, isAllowedInIdentifier(r), true)
	}

	for _, r := range notAllowed {
		assert.Equal(t, isAllowedInIdentifier(r), false)
	}
}

func TestLexNumber(t *testing.T) {
	assert := assert.New(t)

	testLexState(t, "24 ", lexNumber, func(l *Lexer, tokens []*Token) {
		assert.Equal(1, len(tokens))
		assert.Equal(TokenInt, tokens[0].Type)
		assert.Equal("24", tokens[0].Value)
	})

	testLexState(t, "24.56 ", lexNumber, func(l *Lexer, tokens []*Token) {
		assert.Equal(1, len(tokens))
		assert.Equal(TokenFloat, tokens[0].Type)
		assert.Equal("24.56", tokens[0].Value)
	})
}

const testNumRange = `2..5`

func TestLexNumRange(t *testing.T) {
	testLex(t, testNumRange, []expectedToken{
		{"2", TokenInt},
		{"..", TokenRange},
		{"5", TokenInt},
	})
}

const testRecord = `
type alias Foo = 
	{ myInt : Int 
	, myFloat : Float
	}
`

func TestLexRecord(t *testing.T) {
	testLex(t, testRecord, []expectedToken{
		{"type", TokenKeyword},
		{"alias", TokenKeyword},
		{"Foo", TokenIdentifier},
		{"=", TokenAssign},
		{"{", TokenLeftBrace},
		{"myInt", TokenIdentifier},
		{":", TokenColon},
		{"Int", TokenIdentifier},
		{",", TokenComma},
		{"myFloat", TokenIdentifier},
		{":", TokenColon},
		{"Float", TokenIdentifier},
		{"}", TokenRightBrace},
	})
}

const textFuncDecl = `
foo : (Int -> Int) -> Int -> Int
foo fn n =
	fn n
`

func TestLexFuncDecl(t *testing.T) {
	testLex(t, textFuncDecl, []expectedToken{
		{"foo", TokenIdentifier},
		{":", TokenColon},
		{"(", TokenLeftParen},
		{"Int", TokenIdentifier},
		{"->", TokenArrow},
		{"Int", TokenIdentifier},
		{")", TokenRightParen},
		{"->", TokenArrow},
		{"Int", TokenIdentifier},
		{"->", TokenArrow},
		{"Int", TokenIdentifier},
		{"foo", TokenIdentifier},
		{"fn", TokenIdentifier},
		{"n", TokenIdentifier},
		{"=", TokenAssign},
		{"fn", TokenIdentifier},
		{"n", TokenIdentifier},
	})
}

const testRecordUpdate = `
{ model | foo = True }
`

func TestLexRecordUpdate(t *testing.T) {
	testLex(t, testRecordUpdate, []expectedToken{
		{"{", TokenLeftBrace},
		{"model", TokenIdentifier},
		{"|", TokenPipe},
		{"foo", TokenIdentifier},
		{"=", TokenAssign},
		{"True", TokenBool},
		{"}", TokenRightBrace},
	})
}

const testSumType = `
type Op
	= Sum
	| Div
	| Mul
	| Sub
`

func TestSumType(t *testing.T) {
	testLex(t, testSumType, []expectedToken{
		{"type", TokenKeyword},
		{"Op", TokenIdentifier},
		{"=", TokenAssign},
		{"Sum", TokenIdentifier},
		{"|", TokenPipe},
		{"Div", TokenIdentifier},
		{"|", TokenPipe},
		{"Mul", TokenIdentifier},
		{"|", TokenPipe},
		{"Sub", TokenIdentifier},
	})
}

const testString = `
tom = { name = "Tom", bar = "\t\"" }
`

func TestString(t *testing.T) {
	testLex(t, testString, []expectedToken{
		{"tom", TokenIdentifier},
		{"=", TokenAssign},
		{"{", TokenLeftBrace},
		{"name", TokenIdentifier},
		{"=", TokenAssign},
		{`"Tom"`, TokenString},
		{",", TokenComma},
		{"bar", TokenIdentifier},
		{"=", TokenAssign},
		{`"\t\""`, TokenString},
		{"}", TokenRightBrace},
	})
}

const testChar = `
tom = { initial = 'T', foo = '\\' }
`

func TestChar(t *testing.T) {
	testLex(t, testChar, []expectedToken{
		{"tom", TokenIdentifier},
		{"=", TokenAssign},
		{"{", TokenLeftBrace},
		{"initial", TokenIdentifier},
		{"=", TokenAssign},
		{`'T'`, TokenChar},
		{",", TokenComma},
		{"foo", TokenIdentifier},
		{"=", TokenAssign},
		{`'\\'`, TokenChar},
		{"}", TokenRightBrace},
	})
}

const testComment = `
-- comment
-- other comment
`

func TestComment(t *testing.T) {
	testLex(t, testComment, []expectedToken{
		{"-- comment", TokenComment},
		{"-- other comment", TokenComment},
	})
}

const testMultiLineComment = `
{-|-}
{-| Extract the first element of a list.
    head [1,2,3] == Just 1
    head [] == Nothing
-}
`

func TestMultiLineComment(t *testing.T) {
	testLex(t, testMultiLineComment, []expectedToken{
		{"{-|-}", TokenComment},
		{`{-| Extract the first element of a list.
    head [1,2,3] == Just 1
    head [] == Nothing
-}`, TokenComment},
	})
}

const testInfixOp = "theMax = 3 `max` 5"

func TestInfixOp(t *testing.T) {
	testLex(t, testInfixOp, []expectedToken{
		{"theMax", TokenIdentifier},
		{"=", TokenAssign},
		{"3", TokenInt},
		{"`max`", TokenInfixOp},
		{"5", TokenInt},
	})
}

const testList = `
List.map fn [1, 2, 3]
`

func TestList(t *testing.T) {
	testLex(t, testList, []expectedToken{
		{"List", TokenIdentifier},
		{".", TokenDot},
		{"map", TokenIdentifier},
		{"fn", TokenIdentifier},
		{"[", TokenLeftBracket},
		{"1", TokenInt},
		{",", TokenComma},
		{"2", TokenInt},
		{",", TokenComma},
		{"3", TokenInt},
		{"]", TokenRightBracket},
	})
}

const testOp = `
a = [1] ++ [2]
`

func TestLexOp(t *testing.T) {
	testLex(t, testOp, []expectedToken{
		{"a", TokenIdentifier},
		{"=", TokenAssign},
		{"[", TokenLeftBracket},
		{"1", TokenInt},
		{"]", TokenRightBracket},
		{"++", TokenOp},
		{"[", TokenLeftBracket},
		{"2", TokenInt},
		{"]", TokenRightBracket},
	})
}

const testUnclosedString = `
foo = "unclosed
`

func TestLexUnclosedString(t *testing.T) {
	testLex(t, testUnclosedString, []expectedToken{
		{"foo", TokenIdentifier},
		{"=", TokenAssign},
		{"", TokenError},
	})
}

const testUnclosedChar = `
foo = 'a
`

func TestLexUnclosedChar(t *testing.T) {
	testLex(t, testUnclosedChar, []expectedToken{
		{"foo", TokenIdentifier},
		{"=", TokenAssign},
		{"", TokenError},
	})
}

const testBadNumber = `
foo = 12a4
`

func TestLexBadNumber(t *testing.T) {
	testLex(t, testBadNumber, []expectedToken{
		{"foo", TokenIdentifier},
		{"=", TokenAssign},
		{"", TokenError},
	})
}

const testCustomOp = `
foo = 12 -: 13
`

func TestLexCustomOp(t *testing.T) {
	testLex(t, testCustomOp, []expectedToken{
		{"foo", TokenIdentifier},
		{"=", TokenAssign},
		{"12", TokenInt},
		{"-:", TokenOp},
		{"13", TokenInt},
	})
}

type expectedToken struct {
	value string
	typ   TokenType
}

func testLex(t *testing.T, input string, expected []expectedToken) {
	l := New(strings.NewReader(input))
	go l.Run()

	var tokens []*Token
	for {
		tk, ok := l.Next()
		if !ok {
			break
		}
		tokens = append(tokens, tk)
	}

	assert.Equal(t, len(expected), len(tokens))
	for i := range tokens {
		assert.Equal(t, expected[i].typ, tokens[i].Type)
		if tokens[i].Type != TokenError {
			assert.Equal(t, expected[i].value, tokens[i].Value)
		}
	}
}

func testLexState(t *testing.T, input string, fn stateFunc, testFn func(*Lexer, []*Token)) {
	l := New(strings.NewReader(input))
	var tokens []*Token

	wg := new(sync.WaitGroup)
	go func() {
		wg.Add(1)
		for tk := range l.tokens {
			tokens = append(tokens, tk)
		}
		wg.Done()
	}()

	var err error
	l.state, err = fn(l)
	close(l.tokens)
	wg.Wait()

	if err != nil {
		t.Fatal(err)
	}
	testFn(l, tokens)
}
