package main

import (
    "flag"
    "fmt"
    "strings"
)

type Option struct {
    flags       string
    required    bool
    optional    bool
    withno      bool
    short       string
    long        string
    description string
}

func (o *Option) Name() string {
    r := strings.NewReplacer("--", "", "no-", "")
    return r.Replace(o.long)
}

// Check if `arg` matches the short or long flag
func (o *Option) Is(arg string) bool {
    return arg == o.short || arg == o.long
}

func main() {

    flag.String("name", "some", "use name like...")
    flag.Parse()

    fmt.Println(flag.Args())
    num := 0
    num1 := ^num
    fmt.Println(num1)
}
