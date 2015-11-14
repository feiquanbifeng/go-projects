// Just simple principle
// define interpreter

package compiler

import (
    "fmt"
)

// define token type
type KindType string

const (
    INTEGER KindType = "INTEGER"
    PLUS    KindType = "PLUS"
    SUB     KindType = "SUB"
    EOF     KindType = "EOF"
)

type KindInterface interface {
    Print()
}

type ExprInt struct {
    Kind KindType
    i    int
}

func newExprInt(value int) *ExprInt {
    return &ExprInt{
        Kind: INTEGER,
        i:    value,
    }
}

func (i *ExprInt) Print() {
    fmt.Println(i.Kind)
}

type ExprOperate struct {
    Kind KindType
}

func newExprOperate(k KindType) *ExprOperate {
    return &ExprOperate{
        Kind: k,
    }
}

func (s *ExprOperate) Print() {
    fmt.Println(s.Kind)
}
