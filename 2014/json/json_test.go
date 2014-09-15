package json

import (
    "fmt"
    "testing"
)

func TestNewQuiet(t *testing.T) {
    _, err := NewQuiet("./test.json")
    if err != nil {
        t.Errorf("Error new Quiet")
    }
}

func TestLoad(t *testing.T) {
    q, err := NewQuiet("./test.json")
    if err != nil {
        t.Errorf("Error new Quiet")
    }
    q.Load(func() {
        fmt.Println("load success!")
    })
    if apple, ok := q.Data["apple"]; ok {
        fmt.Println(apple)
    }
}

func TestSave(t *testing.T) {
    q, err := NewQuiet("./test.json")
    if err != nil {
        t.Errorf("Error new Quiet")
    }
    q.Data["home"] = "home"
    q.Data["flag"] = false
    q.Save(func() {
        fmt.Println("save success!")
    })
}
