package parents

import (
    "fmt"
    "testing"
)

func TestParents(t *testing.T) {
    first := Parents()
    fmt.Println(first)
    second := Parents("D:\\software\\golang")
    fmt.Println(second)
    third := Parents("D:\\software\\golang", "win32")
    fmt.Println(third)
    forth := Parents("/opt/software/win", "linux")
    fmt.Println(forth)
}
