// Package token 词法单元
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
