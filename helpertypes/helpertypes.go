// Package helpertypes
package helpertypes

import "monkey-interpreter/token"

type TestOutput struct {
	ExpectedType    token.TokenType
	ExpectedLiteral string
}

type TestCase struct {
	Input    string
	Expected []TestOutput
}
