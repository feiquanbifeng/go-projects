// refer to nodejs path-parents
// just simple implements

package parents

import (
    "os"
    "runtime"
    "path"
    "strings"
    "regexp"
    "array"
)

// Public method to provider API
// Actually you could not pass method's parameters
func Parents(args ...string) []string {
    var (
        isWindows = runtime.GOOS == "windows"
        reg *regexp.Regexp
        init array.Array
        cwd string
        sep string
        c string
    )
    if n := len(args); n == 0 {
        cwd, _ = os.Getwd()
    } else if n == 1 {
        cwd = args[0]
    } else {
        cwd = args[0]
        isWindows = strings.HasPrefix(args[1], "win")
    }

    if isWindows {
        c = `[\\\/]`
        init = array.Array{""}
    } else {
        c = `/`
        init = array.Array{"/"}
    }

    reg = regexp.MustCompile(c)

    var join = func(x, y interface{}) array.Array {
        tmpArray := array.Array{x, y}
        var ps = tmpArray.Filter(func(p interface{}, args ...interface{}) bool {
            switch p.(type) {
            case string:
                if p.(string) != "" {
                    return true
                }
                return false
            default:
                return false
            }
            return false
        })
        if isWindows {
            sep = "\\"
        } else {
            sep = "/"
        }
        return array.Array{path.Clean(ps.Join(sep))}
    }

    var res = path.Clean(cwd)
    arr := array.Array{}
    for _, v := range reg.Split(res, -1) {
        arr.Push(v)
    }

    arrReduce := arr.Reduce(func(acc, dir interface{}, ix ...interface{}) interface{} {
        tmpAcc := acc.(array.Array)
        index := ix[0].(int)
        tmpAcc = tmpAcc.Concat(join(tmpAcc[index], dir))
        return tmpAcc
    }, init)

    // Not like javascript support method links
    // everytime you should assign
    arrSlice := arrReduce.(array.Array)
    arrReverse := arrSlice.Slice(1, 0)
    arrReverse.Reverse()
    if len(arrReverse) >= 2 {
        if arrReverse[0] == arrReverse[1] {
            return []string{arrReverse[0].(string)}
        }
    }

    if isWindows && strings.HasPrefix(cwd, "\\") {
        cut := arrReverse.Slice(0, -1)
        cut.Map(func(d interface{}, args ...interface{}) interface{} {
            var ch = d.(string)[0]
            if ch == '\\' {
                return d
            } else if ch == '.' {
                return "\\" + d.(string)[1:]
            } else {
                return "\\" + d.(string)
            }
        })
        return cut.ToString()
    }
    return arrReverse.ToString()
}
