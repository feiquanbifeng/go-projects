// refer to
// 

package main

import (
	"os"
	"runtime"
	"fmt"
	"path"
	"strings"
	"regexp"
	"array"
)

func Parents(cwd string, opts ...string) []string {
	if cwd == "" {
		cwd, _ = os.Getwd()
	}
	var (
		isWindows bool
		sep string
		reg *regexp.Regexp
		c string
		init array.Array
	)
	if len(opts) == 0 {
		isWindows = runtime.GOOS == "windows"
	} else {
		isWindows = strings.HasPrefix(opts[0], "win")
	}
	if isWindows {
		c = `[\\\/]`
		init = array.Array{""}
	} else {
		c = `/`
		init = array.Array{"/"}
	}

	reg, _ = regexp.Compile(c)

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
	rr := array.Array{"\\F:", "Go", "1110"}
	reg.Split(res, 5)
	arrReduce := rr.Reduce(func(acc, dir interface{}, ix ...interface{}) interface{} {
		tmpAcc := acc.(array.Array)
		index := ix[0].(int)
		tmpAcc = tmpAcc.Concat(join(tmpAcc[index], dir))
		return tmpAcc
	}, init)

	arrSlice := arrReduce.(array.Array)
	arrReverse := arrSlice.Slice(1, 0)
	arrReverse.Reverse()
	fmt.Println(arrReverse)

	if arrReverse[0] == arrReverse[1] {
		return []string{arrReverse[0].(string)}
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
		fmt.Println(([]string)(cut))
	}
	return nil
}

func main() {
	fmt.Println(Parents("\\F:\\Go\\1110"))
}
