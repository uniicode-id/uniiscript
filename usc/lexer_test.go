package usc_test

import (
	"testing"

	"github.com/uniicode-id/uniiscript/usc"
)

func TestOneLine(t *testing.T) {
	tests := []struct {
		source string
		tokens []usc.TokenKind
	}{
		{
			source: "var a = 1;",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenInt,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
		{
			source: "var a = 1; println(a);",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenInt,
				usc.TokenSemicolon,
				usc.TokenIdent,
				usc.TokenLeftParen,
				usc.TokenIdent,
				usc.TokenRightParen,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
		{
			source: "var s = \"Hello!\\nMy name is Adam!\";",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenStr,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
		{
			source: "var c = 'a'; // single",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenChar,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
		{
			source: " /* multi line */ var c = 'a'; // single",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenChar,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
		{
			source: "var f = 1.0;",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenFloat,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
		{
			source: "var n = 100_000;",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenInt,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
		{
			source: "var f2 = 10_0.00_0;",
			tokens: []usc.TokenKind{
				usc.TokenVar,
				usc.TokenIdent,
				usc.TokenAssign,
				usc.TokenFloat,
				usc.TokenSemicolon,
				usc.TokenEOF,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			l, err := usc.NewLexer("unknown", []byte(tt.source), nil)
			if err != nil {
				t.Fatal(err)
			}

			for _, want := range tt.tokens {
				l = l.Next()

				if got := l.Tok.Kind; got != want {
					t.Fatalf("want %v, got %v", want, got)
				}
			}
		})
	}
}

func TestChar(t *testing.T) {
	tests := []struct {
		source string
		char   byte
	}{
		{
			source: "'a'",
			char:   'a',
		},
		{
			source: "'\\n'",
			char:   '\n',
		},
	}

	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			l, err := usc.NewLexer("unknown", []byte(tt.source), nil)
			if err != nil {
				t.Fatal(err)
			}

			l = l.Next()

			if got := l.Tok.Kind; got != usc.TokenChar {
				t.Fatalf("want %v, got %v", usc.TokenChar, got)
			}

			if got := l.Tok.Value.(byte); got != tt.char {
				t.Fatalf("want %v, got %v", tt.char, got)
			}
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		source   string
		expected int64
	}{
		{
			source:   "20",
			expected: 20,
		},
		{
			source:   "00",
			expected: 0,
		},
		{
			source:   "30_0",
			expected: 300,
		},
	}

	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			l, err := usc.NewLexer("unknown", []byte(tt.source), nil)
			if err != nil {
				t.Fatal(err)
			}

			l = l.Next()

			if got := l.Tok.Kind; got != usc.TokenInt {
				t.Fatalf("want %v, got %v", usc.TokenInt, got)
			}

			if got := l.Tok.Value; got != tt.expected {
				t.Fatalf("want %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestFloat(t *testing.T) {
	tests := []struct {
		source   string
		expected float64
	}{
		{
			source:   "2.0",
			expected: 2.0,
		},
		{
			source:   "0.0",
			expected: 0.0,
		},
		{
			source:   "30_0.0",
			expected: 300.0,
		},
		{
			source:   "30_0.43_20",
			expected: 300.432,
		},
	}

	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			l, err := usc.NewLexer("unknown", []byte(tt.source), nil)
			if err != nil {
				t.Fatal(err)
			}

			l = l.Next()

			if got := l.Tok.Kind; got != usc.TokenFloat {
				t.Fatalf("want %v, got %v", usc.TokenFloat, got)
			}

			if got := l.Tok.Value; got != tt.expected {
				t.Fatalf("want %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		source   string
		expected string
	}{
		{
			source:   "\"hello world\"",
			expected: "hello world",
		},
		{
			source:   "\"hello\\nworld\"",
			expected: "hello\nworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			l, err := usc.NewLexer("unknown", []byte(tt.source), nil)
			if err != nil {
				t.Fatal(err)
			}

			l = l.Next()

			if got := l.Tok.Kind; got != usc.TokenStr {
				t.Fatalf("want %v, got %v", usc.TokenStr, got)
			}

			if got := l.Tok.Value; got != tt.expected {
				t.Fatalf("want %v, got %v", tt.expected, got)
			}
		})
	}
}
