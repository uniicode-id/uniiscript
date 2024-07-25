package usc

import (
	"strconv"

	"github.com/uniicode-id/uniiscript/common"
)

type Lexer struct {
	SourceName       string
	Source           []byte
	Current          byte
	IsEscaped        bool
	Index, Line, Pos uint
	Tok, PrevTok     Token

	errorCallback ErrorCallback
}

func NewLexer(
	sourceName string,
	source []byte,
	errorCallback ErrorCallback,
) (Lexer, error) {
	return Lexer{
		SourceName:    sourceName,
		Source:        []byte(source),
		Index:         0,
		Line:          1,
		Pos:           1,
		Tok:           Token{},
		PrevTok:       Token{},
		errorCallback: errorCallback,
	}.read(), nil
}

func (l Lexer) read() Lexer {
	l.IsEscaped = false

	if l.Index >= uint(len(l.Source)) {
		l.Current = 0
		return l
	}
	l.Current = l.Source[l.Index]
	l.Index++
	l.Pos++

	return l
}

func (l Lexer) readEscape() Lexer {
	l = l.read()
	b := l.Current

	if b == '\\' {
		l = l.read()
		b = l.Current

		switch b {
		case 'n':
			l.Current = '\n'
		case 'r':
			l.Current = '\r'
		case 't':
			l.Current = '\t'
		case '\\':
			l.Current = '\\'
		case '\'':
			l.Current = '\''
		case '"':
			l.Current = '"'
		case '0':
			l.Current = 0
		default:
			l.errorCallback("Invalid escape sequence")
			l.Tok.Kind = TokenIllegal
			return l
		}

		l.IsEscaped = true
	}

	return l
}

func (l Lexer) peek() byte {
	if l.Index < uint(len(l.Source)) {
		return l.Source[l.Index]
	}

	return 0
}

func (l Lexer) peekDouble() byte {
	if l.Index+1 < uint(len(l.Source)) {
		return l.Source[l.Index+1]
	}

	return 0

}

func (l Lexer) comment() Lexer {
	b := l.Current
	p := l.peek()

	if b == '/' {
		if p == '/' {
			l = l.read()
			b = l.Current

			for b != 0 && b != '\n' {
				l = l.read()
				b = l.Current
			}

			l = l.read()
		} else if p == '*' {
			l = l.read()
			b = l.Current
			p = l.peek()

			for b != 0 && (b != '*' || p != '/') {
				l = l.read()
				b = l.Current
				p = l.peek()
			}
			l = l.read().read()
		}
	}

	return l
}

func (l Lexer) whitespaceComment() Lexer {
	b := l.Current

	for b == ' ' || b == '\t' || b == '\n' || b == '\r' || b == '/' {
		if b == '/' {
			l = l.comment()
		} else {
			l = l.read()
		}

		b = l.Current
	}

	return l
}

func (l Lexer) ident() Lexer {
	l.Tok.Kind = TokenIdent

	b := l.Current

	value := TokenValueIdent{}

	isFirst := true

	for (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || b == '_' || (b >= '0' && b <= '9' && !isFirst) {
		isFirst = false

		value.Name += string(b)

		l = l.read()
		b = l.Current
	}

	value.Hash = common.Hash(value.Name)

	l.Tok.Value = value

	keywordToken, isKeyword := keywords[value.Name]
	if isKeyword {
		l.Tok.Kind = keywordToken
	}

	return l
}

func (l Lexer) number() Lexer {
	l.Tok.Kind = TokenInt

	b := l.Current

	value := ""

	for (b >= '0' && b <= '9') || b == '_' {
		value += string(b)

		l = l.read()
		b = l.Current
	}

	var err error
	l.Tok.Value, err = strconv.ParseInt(value, 0, 64)
	if err != nil {
		l.errorCallback("Invalid integer literal")
		l.Tok.Kind = TokenIllegal
		return l
	}

	if b == '.' {
		l.Tok.Kind = TokenFloat

		value += "."

		l = l.read()
		b = l.Current

		for (b >= '0' && b <= '9') || b == '_' {
			value += string(b)

			l = l.read()
			b = l.Current
		}

		l.Tok.Value, err = strconv.ParseFloat(value, 64)
		if err != nil {
			l.errorCallback("Invalid float literal")
			l.Tok.Kind = TokenIllegal
			return l
		}
	}

	return l
}

func (l Lexer) char() Lexer {
	l.Tok.Kind = TokenChar

	b := l.Current
	if b != '\'' {
		l.errorCallback("Invalid character literal")
		l.Tok.Kind = TokenIllegal
		return l
	}

	l = l.readEscape()
	l.Tok.Value = l.Current

	l = l.read()
	b = l.Current
	if b != '\'' {
		l.errorCallback("Invalid character literal")
		l.Tok.Kind = TokenIllegal
		return l
	}

	return l.read()
}

func (l Lexer) string() Lexer {
	l.Tok.Kind = TokenStr

	b := l.Current
	if b != '"' {
		l.errorCallback("Invalid string literal")
		l.Tok.Kind = TokenIllegal
		return l
	}

	value := ""

	l = l.read()
	b = l.Current

	for b != 0 && b != '"' {
		value += string(b)

		l = l.readEscape()
		b = l.Current

		if b == '\n' && !l.IsEscaped {
			l.errorCallback("Invalid string literal")
			l.Tok.Kind = TokenIllegal
			return l
		}
	}

	if b == 0 {
		l.errorCallback("Invalid string literal")
		l.Tok.Kind = TokenIllegal
		return l
	}

	l.Tok.Value = value

	l = l.read()

	return l
}

func (l Lexer) operator() Lexer {
	l.Tok.Kind = TokenIllegal

	b := l.Current

	switch b {
	case '+':
		l.Tok.Kind = TokenAdd
		p := l.peek()
		switch p {
		case '+':
			l = l.read()
			l.Tok.Kind = TokenInc
		case '=':
			l = l.read()
			l.Tok.Kind = TokenAddAssign
		}
	case '-':
		l.Tok.Kind = TokenSub
		p := l.peek()
		switch p {
		case '-':
			l = l.read()
			l.Tok.Kind = TokenDec
		case '=':
			l = l.read()
			l.Tok.Kind = TokenSubAssign
		}
	case '*':
		l.Tok.Kind = TokenMul
		p := l.peek()
		switch p {
		case '*':
			l = l.read()
			l.Tok.Kind = TokenPow
			p = l.peek()
			switch p {
			case '=':
				l = l.read()
				l.Tok.Kind = TokenPowAssign
			}
		case '=':
			l = l.read()
			l.Tok.Kind = TokenMulAssign
		}
	case '/':
		l.Tok.Kind = TokenDiv
		p := l.peek()
		switch p {
		case '=':
			l = l.read()
			l.Tok.Kind = TokenDivAssign
		}
	case '%':
		l.Tok.Kind = TokenMod
		p := l.peek()
		switch p {
		case '=':
			l = l.read()
			l.Tok.Kind = TokenModAssign
		}
	case '&':
		l.Tok.Kind = TokenBitAnd
		p := l.peek()
		switch p {
		case '&':
			l = l.read()
			l.Tok.Kind = TokenAnd
		case '=':
			l = l.read()
			l.Tok.Kind = TokenBitAndAssign
		}
	case '|':
		l.Tok.Kind = TokenBitOr
		p := l.peek()
		switch p {
		case '|':
			l = l.read()
			l.Tok.Kind = TokenOr
		case '=':
			l = l.read()
			l.Tok.Kind = TokenBitOrAssign
		}
	case '^':
		l.Tok.Kind = TokenBitXor
		p := l.peek()
		switch p {
		case '^':
			l = l.read()
			l.Tok.Kind = TokenXor
		case '=':
			l = l.read()
			l.Tok.Kind = TokenBitXorAssign
		}
	case '~':
		l.Tok.Kind = TokenBitNot
	case '<':
		l.Tok.Kind = TokenLess
		p := l.peek()
		switch p {
		case '<':
			l = l.read()
			l.Tok.Kind = TokenShiftLeft
			p = l.peek()
			switch p {
			case '=':
				l = l.read()
				l.Tok.Kind = TokenShiftLeftAssign
			}
		case '=':
			l = l.read()
			l.Tok.Kind = TokenLessEqual
		}
	case '>':
		l.Tok.Kind = TokenGreater
		p := l.peek()
		switch p {
		case '>':
			l = l.read()
			l.Tok.Kind = TokenShiftRight
			p = l.peek()
			switch p {
			case '=':
				l = l.read()
				l.Tok.Kind = TokenShiftRightAssign
			}
		case '=':
			l = l.read()
			l.Tok.Kind = TokenGreaterEqual
		}
	case '!':
		l.Tok.Kind = TokenNot
		p := l.peek()
		switch p {
		case '=':
			l = l.read()
			l.Tok.Kind = TokenNotEqual
		}
	case '=':
		l.Tok.Kind = TokenAssign
		p := l.peek()
		switch p {
		case '=':
			l = l.read()
			l.Tok.Kind = TokenEqual
		}
	case '?':
		l.Tok.Kind = TokenQuestion
	case '(':
		l.Tok.Kind = TokenLeftParen
	case ')':
		l.Tok.Kind = TokenRightParen
	case '{':
		l.Tok.Kind = TokenLeftBrace
	case '}':
		l.Tok.Kind = TokenRightBrace
	case '[':
		l.Tok.Kind = TokenLeftBracket
	case ']':
		l.Tok.Kind = TokenRightBracket
	case ',':
		l.Tok.Kind = TokenComma
	case ':':
		l.Tok.Kind = TokenColon
	case ';':
		l.Tok.Kind = TokenSemicolon
	case '.':
		l.Tok.Kind = TokenDot
		p := l.peek()
		pd := l.peekDouble()
		if p == '.' && pd == '.' {
			l = l.read().read()
			l.Tok.Kind = TokenEllipsis
		}
	}

	l = l.read()

	return l
}

func (l Lexer) Next() Lexer {
	l.Tok.Kind = TokenIllegal

	l = l.whitespaceComment()

	b := l.Current
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || b == '_' {
		l = l.ident()
	} else if b >= '0' && b <= '9' || b == '.' {
		l = l.number()
	} else if b == '\'' {
		l = l.char()
	} else if b == '"' {
		l = l.string()
	} else if b == 0 {
		l.Tok.Kind = TokenEOF
	} else {
		l = l.operator()
	}

	if l.Tok.Kind != TokenIllegal {
		l.PrevTok = l.Tok
	} else {
		l.errorCallback("Illegal token")
	}

	return l
}
