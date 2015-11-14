// Just simple principle
// define interpreter

package compiler

import (
    "testing"
)

func TestExpr(t *testing.T) {
    text := "file.txt"
    interpreter := &Interpreter{
        text:  text,
        bytes: []byte{},
        pos:   0,
        tree:  nil,
    }
    interpreter.Expr()
}
