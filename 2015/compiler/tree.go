// Just simple complier
// define interpreter

package compiler

import ()

const (
    cReg string = `[^0][0-9]*`
    oReg string = `[\+|\-|\*|\%]{1}`
)

// All tree
type Tree struct {
    left  *Tree
    right *Tree
    value KindInterface
}

func newRootTree(v KindInterface) *Tree {
    return &Tree{
        left:  nil,
        right: nil,
        value: v,
    }
}

// Add left node
func (t *Tree) AppendLeft(v KindInterface) {
    temp := t
    t.left = &Tree{
        left:  temp.left,
        right: temp.right,
        value: temp.value,
    }
    t.value = v
}

// Add right node
func (t *Tree) AppendRight(v KindInterface) {
    t.right = &Tree{
        left:  nil,
        right: nil,
        value: v,
    }
}
