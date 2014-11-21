// vary just like npmjs vary
// support two public method

package vary

import (
    "net/http"
    "strings"
    "regexp"
)

// Public method Append
// the second parameter is an array
func Append(header string, argv ...string) string {
    var (
        greg = regexp.MustCompile(`[\(\)<>@,;:\\"\/\[\]\?=\{\}\s\t]`)
        fields []string
    )
    if len(argv) == 1 {
        fields = append(fields, parse(argv[0])...)
    } else {
        fields = append(fields, argv...)
    }
    for _, v := range fields {
        if greg.MatchString(v) {
            panic("field argument contains an invalid header")
        }
    }
    // existing, unspecified vary
    if header == "*" {
        return header
    }
    // enumerate current values
    vals := parse(strings.ToLower(header))

    for _, v := range fields {
        if v == "*" {
            return "*"
        }
        field := strings.ToLower(v)
        var flag bool
        for _, k := range vals {
            if k == "*" {
                return "*"
            }
            if k != field {
                flag = true
            }
        }
        if flag {
            vals = append(vals, field)
            if header != "" {
                header += ", " + v
            } else {
                header = v
            }
        }
    }
    return header
}

// Parse string throw `,`
func parse(header string) []string {
    reg := regexp.MustCompile(`/ *, */`)
    return reg.Split(strings.TrimSpace(header), -1)
}

// Public method
func Vary(res *http.Response, argv ...string) {
    head := res.Header[http.CanonicalHeaderKey("Vary")]
    header := strings.Join(head, ",")
    res.Header.Set("Vary", Append(header, argv...))
}
