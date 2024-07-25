package usc

type Compiler struct {
	Lexer Lexer
}

func NewCompiler() Compiler {
	return Compiler{}
}
