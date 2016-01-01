package utils

import (
    "fmt"
    "testing"
)

func TestTripleDESEncode(t *testing.T) {
    encode := TripleDESEncode([]byte("2"))
    fmt.Println(encode)
    decode := TripleDESDecode(encode)
    fmt.Println(decode)
}
