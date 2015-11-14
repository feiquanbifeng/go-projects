// Just simple principle
// define interpreter

package compiler

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strconv"
)

type Interpreter struct {
    // input text
    text string
    // bytes
    bytes []byte
    // current index
    pos int
    // root tree
    tree *Tree
}

func (i *Interpreter) Expr() {

    // load from a file
    // begin to explain the record
    bytes, err := ioutil.ReadFile(i.text)
    if err != nil {
        panic(err)
    }
    i.bytes = bytes

    i.RegExpInt()
    i.RegExpOperate()
    i.RegExpInt()

    var result int = 0
    var value KindInterface

    if i.tree != nil {

        l := i.tree.left
        if l != nil {
            value = l.value
            if c, ok := value.(*ExprInt); ok {
                result += c.i
            }
        }

        r := i.tree.right
        if r != nil {
            value = r.value
            if c, ok := value.(*ExprInt); ok {

                op := i.tree.value
                if o, ok := op.(*ExprOperate); ok {
                    switch o.Kind {
                    case PLUS:
                        result += c.i
                    case SUB:
                        result -= c.i
                    default:
                    }
                }
            }
        }
    }
    fmt.Println("The result is :", result)
}

func (i *Interpreter) RegExpInt() {

    var index []int
    var bytes []byte = i.bytes
    // if match number first time
    reg := regexp.MustCompile(cReg)
    index = reg.FindIndex(bytes)

    // if index is not nil then init root tree
    if index != nil {

        v, err := strconv.Atoi((string)(bytes[index[0]:index[1]]))

        if err != nil {
            panic(err)
        }

        if i.tree == nil {
            i.tree = newRootTree(newExprInt(v))
        } else {
            i.tree.AppendRight(newExprInt(v))
        }

        pos := index[1]
        i.bytes = bytes[pos:]
        i.pos += pos
    }
}

func (i *Interpreter) RegExpOperate() {

    var index []int
    var bytes []byte = i.bytes
    // if match operater first time
    reg := regexp.MustCompile(oReg)
    index = reg.FindIndex(bytes)

    // if index is not nil then set the root
    if index != nil {
        v := (string)(bytes[index[0]:index[1]])
        if v == "+" {
            i.tree.AppendLeft(newExprOperate(PLUS))
        } else if v == "-" {
            i.tree.AppendLeft(newExprOperate(SUB))
        }
        pos := index[1]
        i.bytes = bytes[pos:]
        i.pos += pos
    }

}
