package main

import (
    "flag"
    "fmt"
    "math"
    "os"
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

type Command struct {
    commands []string
    options  []Option
    execs    []string
    args     []string
    name     string
    version  string
}

func (c *Command) command(name, desc string) {
    args := strings.Fields(name)
    // cmd := &
}

func (c *Command) addImplicitHelpCommand() {
    c.command("help [cmd]", "display help for [cmd]")
}

func (c *Command) parseExpectedArgs(args string) {
}

// Return an option matching `arg` if any
func (c *Command) OptionFor(arg) Option {
    for _, v := range c.options {
        if v.Is(arg) {
            return v
        }
    }
    return nil
}

// Return an object containing options as key-vale pairs
func (c *Command) Opts() map[string]string {
    result := make(map[string]string, len(c.options))
    for _, v := range c.options {
        key := v.Name()
        if key == "version" {
            result[key] = c.version
        } else {
            // result[key] =
        }
    }
    return result
}

// Unknown option `flag`
func (c *Command) unknownOption(flag string) {
    fmt.Fprintf(os.Stderr, " error: unknown option %s", flag)
    os.Exit(1)
}

// Set the program version to `str`.
func (c *Command) Version(args ...string) string {
    if len(args) == 0 {
        return c.version
    }
    c.version = args[0]
    flags := "-V, --version"
    if len(args) == 2 {
        flags = args[1]
    }
}

// Set the description to `str`.
func (c *Command) Description(args ...string) string {
    if len(args) == 0 {
        return ""
    }
    return ""
}

// Set an alias for the command
func (c *Command) Alias(args ...string) string {
    if len(args) == 0 {
        return ""
    }
    return ""
}

// Set / Get the command usage `str`
func (c *Command) Usage(str string) {

}

// Return program help document.
func (c *Command) HelpInformation() string {
    msg := []string {
        "",
        " Usage: " + c.name,
        c.commandHelp(),
        " Options:",
        "",
        c.optionHelp(),
        "",
        ""
    }
    return strings.Join(msg, "\n")
}

// Output help information for this command
func (c *Command) OutputHelp() {
    fmt.Fprintf(os.Stdout, c.HelpInfomation())
}

// Output help information and exit.
func (c *Command) Help() {
    c.OutputHelp()
    os.Exit(0)
}

// Pad `str` to `witdh`.
func Pad(str string, width int) string {
    count := math.Max(0, float64(width-len(str)))
    return str + strings.Repeat(" ", int(count)+1)
}

// Output help information if necessary
func OutputHelpIfNecessary(cmd Commander, options []Option) {
    for _, v := range options {
        if v == "--help" || v == "-h" {
            cmd.outputHelp()
            os.Exit(0)
        }
    }
}

func main() {

    flag.String("name", "some", "use name like...")
    flag.Parse()

    fmt.Println(flag.Args())
    num := 0
    num1 := ^num
    fmt.Println(num1)
}
