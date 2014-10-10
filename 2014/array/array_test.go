package array

import (
    "fmt"
    "testing"
)

var arr = Array{12, 5, 8, 7, 4}

func TestLength(t *testing.T) {
    if arr.Length() == 5 {
        fmt.Println("length equal!")
    }
}

func TestShift(t *testing.T) {
    first := arr.Shift()
    if first != 12 {
        t.Errorf("Shift method error")
    }
}
