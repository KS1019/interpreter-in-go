package parser

import (
	"interpreter-in-go/ast"
	"interpreter-in-go/lexer"
	"interpreter-in-go/token"
)

type Paser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Paser {
	p := &Paser{l: l}

	//二つのトークンを読み込む
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Paser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Paser) ParseProgram() *ast.Program {
	return nil
}
