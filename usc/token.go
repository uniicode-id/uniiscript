package usc

import "github.com/uniicode-id/uniiscript/common"

type TokenKind uint8

const (
	TokenIllegal TokenKind = iota

	TokenVar
	TokenConst
	TokenFunc
	TokenReturn
	TokenIf
	TokenElse
	TokenSwitch
	TokenCase
	TokenDefault
	TokenFallthrough
	TokenWhile
	TokenDo
	TokenFor
	TokenBreak
	TokenContinue
	TokenGoto
	TokenLabel
	TokenClass
	TokenAbstract
	TokenInterface
	TokenEnum
	TokenStatic
	TokenPublic
	TokenPrivate
	TokenProtected
	TokenImport
	TokenExport

	TokenAdd
	TokenSub
	TokenMul
	TokenDiv
	TokenMod
	TokenPow
	TokenAddAssign
	TokenSubAssign
	TokenMulAssign
	TokenDivAssign
	TokenModAssign
	TokenPowAssign

	TokenInc
	TokenDec

	TokenBitAnd
	TokenBitOr
	TokenBitXor
	TokenBitNot
	TokenShiftLeft
	TokenShiftRight
	TokenBitAndAssign
	TokenBitOrAssign
	TokenBitXorAssign
	TokenShiftLeftAssign
	TokenShiftRightAssign

	TokenAnd
	TokenOr
	TokenXor
	TokenNot

	TokenEqual
	TokenNotEqual
	TokenLess
	TokenLessEqual
	TokenGreater
	TokenGreaterEqual

	TokenQuestion

	TokenAssign

	TokenLeftParen
	TokenRightParen
	TokenLeftBrace
	TokenRightBrace
	TokenLeftBracket
	TokenRightBracket

	TokenComma
	TokenColon
	TokenSemicolon
	TokenDot
	TokenEllipsis

	TokenIdent

	TokenStr
	TokenChar
	TokenInt
	TokenFloat

	TokenEOF
)

func (kind TokenKind) String() string {
	switch kind {
	case TokenVar:
		return "var"
	case TokenConst:
		return "const"
	case TokenFunc:
		return "func"
	case TokenReturn:
		return "return"
	case TokenIf:
		return "if"
	case TokenElse:
		return "else"
	case TokenSwitch:
		return "switch"
	case TokenCase:
		return "case"
	case TokenDefault:
		return "default"
	case TokenFallthrough:
		return "fallthrough"
	case TokenWhile:
		return "while"
	case TokenDo:
		return "do"
	case TokenFor:
		return "for"
	case TokenBreak:
		return "break"
	case TokenContinue:
		return "continue"
	case TokenGoto:
		return "goto"
	case TokenLabel:
		return "label"
	case TokenClass:
		return "class"
	case TokenAbstract:
		return "abstract"
	case TokenInterface:
		return "interface"
	case TokenEnum:
		return "enum"
	case TokenStatic:
		return "static"
	case TokenPublic:
		return "public"
	case TokenPrivate:
		return "private"
	case TokenProtected:
		return "protected"
	case TokenImport:
		return "import"
	case TokenExport:
		return "export"
	case TokenAdd:
		return "+"
	case TokenSub:
		return "-"
	case TokenMul:
		return "*"
	case TokenDiv:
		return "/"
	case TokenMod:
		return "%"
	case TokenPow:
		return "**"
	case TokenAddAssign:
		return "+="
	case TokenSubAssign:
		return "-="
	case TokenMulAssign:
		return "*="
	case TokenDivAssign:
		return "/="
	case TokenModAssign:
		return "%="
	case TokenPowAssign:
		return "**="
	case TokenInc:
		return "++"
	case TokenDec:
		return "--"
	case TokenBitAnd:
		return "&"
	case TokenBitOr:
		return "|"
	case TokenBitXor:
		return "^"
	case TokenBitNot:
		return "~"
	case TokenShiftLeft:
		return "<<"
	case TokenShiftRight:
		return ">>"
	case TokenBitAndAssign:
		return "&="
	case TokenBitOrAssign:
		return "|="
	case TokenBitXorAssign:
		return "^="
	case TokenShiftLeftAssign:
		return "<<="
	case TokenShiftRightAssign:
		return ">>="
	case TokenAnd:
		return "&&"
	case TokenOr:
		return "||"
	case TokenXor:
		return "^^"
	case TokenNot:
		return "!"
	case TokenEqual:
		return "=="
	case TokenNotEqual:
		return "!="
	case TokenLess:
		return "<"
	case TokenLessEqual:
		return "<="
	case TokenGreater:
		return ">"
	case TokenGreaterEqual:
		return ">="
	case TokenQuestion:
		return "?"
	case TokenAssign:
		return "="
	case TokenLeftParen:
		return "("
	case TokenRightParen:
		return ")"
	case TokenLeftBrace:
		return "{"
	case TokenRightBrace:
		return "}"
	case TokenLeftBracket:
		return "["
	case TokenRightBracket:
		return "]"
	case TokenComma:
		return ","
	case TokenColon:
		return ":"
	case TokenSemicolon:
		return ";"
	case TokenDot:
		return "."
	case TokenEllipsis:
		return "..."
	case TokenIdent:
		return "ident"
	case TokenStr:
		return "string"
	case TokenChar:
		return "char"
	case TokenInt:
		return "integer"
	case TokenFloat:
		return "float"
	case TokenEOF:
		return "EOF"
	default:
		return "illegal"
	}
}

func (kind TokenKind) Hash() int {
	return common.Hash(kind.String())
}

func (kind TokenKind) IsKeyword() bool {
	return kind >= TokenVar && kind <= TokenExport
}

var keywords = map[string]TokenKind{
	"var":         TokenVar,
	"const":       TokenConst,
	"func":        TokenFunc,
	"return":      TokenReturn,
	"if":          TokenIf,
	"else":        TokenElse,
	"switch":      TokenSwitch,
	"case":        TokenCase,
	"default":     TokenDefault,
	"fallthrough": TokenFallthrough,
	"while":       TokenWhile,
	"do":          TokenDo,
	"for":         TokenFor,
	"break":       TokenBreak,
	"continue":    TokenContinue,
	"goto":        TokenGoto,
	"label":       TokenLabel,
	"class":       TokenClass,
	"abstract":    TokenAbstract,
	"interface":   TokenInterface,
	"enum":        TokenEnum,
	"static":      TokenStatic,
	"public":      TokenPublic,
	"private":     TokenPrivate,
	"protected":   TokenProtected,
	"import":      TokenImport,
	"export":      TokenExport,
}

type Token struct {
	Kind      TokenKind
	Value     any
	Line, Pos uint
}

type TokenValueIdent struct {
	Name string
	Hash int
}
