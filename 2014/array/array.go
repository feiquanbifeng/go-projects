// reference javascript array method
// still some method not implements

package array

import (
    "reflect"
    "strconv"
    "strings"
)

type Array []interface{}

// Return the length of elements.
func (a *Array) Length() int {
    return len(*a)
}

// Remove first element from an array and returns that element.
// this method changes the length of the array.
func (a *Array) Shift() (first interface{}) {
    if a.Length() == 0 {
        return "<nil>"
    }
    first = (*a)[0]
    *a = (*a)[1:]
    return
}

// Adds one or more elements to the beginning of an array.
// this method return the new length of the array.
func (a *Array) Unshift(args ...interface{}) int {
    *a = append(Array(args), *a...)
    return a.Length()
}

// Remove the last element from an array and returns that element.
func (a *Array) Pop() (last interface{}) {
    l := a.Length()
    if l == 0 {
        return "<nil>"
    }
    last = (*a)[l-1]
    *a = (*a)[:l-1]
    return
}

// Adds one or more elements to the end of an array.
// this method return the new length of the array.
func (a *Array) Push(args ...interface{}) int {
    *a = append(*a, args...)
    return a.Length()
}

// Test whether some element in the array passes the test implemented
// by the provided function.
func (a *Array) Some(fn func(element interface{}, args ...interface{}) bool) bool {
    for i, v := range *a {
        if fn(v, i, *a) {
            return true
        }
    }
    return false
}

// Tests whether all elements in the array passes the test implemented
// by the provided function.
func (a *Array) Every(fn func(element interface{}, args ...interface{}) bool) bool {
    for i, v := range *a {
        if !fn(v, i, *a) {
            return false
        }
    }
    return true
}

// Creates a new array with all elements that pass the test implemented by the provided function.
func (a *Array) Filter(fn func(element interface{}, args ...interface{}) bool) Array {
    var tmp Array
    for i, v := range *a {
        if fn(v, i, *a) {
            tmp = append(tmp, v)
        }
    }
    return tmp
}

// Creates a new array with the results of calling a provided function on every element in this array.
func (a *Array) Map(fn func(element interface{}, args ...interface{}) interface{}) Array {
    var tmp Array
    for i, v := range *a {
        tmp = append(tmp, fn(v, i, *a))
    }
    return tmp
}

// Reduce method applies a function against an accumulator and each value of the array
// has to reduce it to a single value
func (a *Array) Reduce(fn func(prev, curr interface{}, args ...interface{}) interface{}, init ...interface{}) interface{} {
    if a.Length() == 0 {
        return nil
    }
    var (
        r interface{}
        i = 0
    )
    if len(init) != 0 {
        r = init[0]
    } else {
        r = (*a)[0]
        i = 1
    }
    for ; i < a.Length(); i++ {
        r = fn(r, (*a)[i])
    }
    return r
}

// Private method translate all element to string.
func convertString(arr *Array) []string {
    return interfaceToString(arr)
}

// Private method translate all element to string. 
func interfaceToString(a *Array) []string {
    str := make([]string, 0, a.Length())
    for _, v := range *a {
        switch v.(type) {
        case byte:
            str = append(str, string(v.(byte)))
        case string:
            str = append(str, v.(string))
        case bool:
            str = append(str, strconv.FormatBool(v.(bool)))
        case int:
            str = append(str, strconv.Itoa(v.(int)))
        case int64:
            str = append(str, strconv.FormatInt(v.(int64), 10))
        case uint64:
            str = append(str, strconv.FormatUint(v.(uint64), 10))
        }
    }
    return str
}

// Convert to string
func (a *Array) ToString() []string {
    return interfaceToString(a)
}

// Joins all elements of an array into a string.
func (a *Array) Join(sep string) string {
    return strings.Join(convertString(a), sep)
}

// Returns a new array comprised of the array on which it is called
// joined with the array(s) and/or value(s) provided as arguments.
func (a *Array) Concat(arr ...[]interface{}) []interface{} {
    var tmp Array
    tmp = append(tmp, *a...)
    for _, v := range arr {
        tmp = append(tmp, v...)
    }
    return tmp
}

// Return true if param is an array, false if it is not.
func (a *Array) IsArray(arg interface{}) bool {
    v := reflect.ValueOf(arg)
    switch v.Kind() {
    case reflect.Slice, reflect.Array:
        return true
    }
    return false
}

// Executes a provided function once per array element.
func (a *Array) ForEach(fn func(element interface{}, args ...interface{})) {
    for i, v := range *a {
        fn(v, i, *a)
    }
}

// Returns the first index at which a given element can be found in the array
// or -1 if it is not present.
func (a *Array) IndexOf(elem interface{}, from int) int {
    if from < 0 {
        from = 0
    }
    for i := from; i < a.Length(); i++ {
        if (*a)[i] == elem {
            return i
        }
    }
    return -1
}

// Reverses an array in place
// the first array element becomes the last and the last becomes the first.
func (a *Array) Reverse() {
    l := a.Length()
    tmp := make([]interface{}, 0, l)
    for i := l - 1; i >= 0; i-- {
        tmp = append(tmp, (*a)[i])
    }
    *a = tmp
}

// Returns a shallow copy of a portion of an array into a new array object.
func (a *Array) Slice(begin, end int) Array {
    l := a.Length()
    if end <= 0 {
        end = end + l
        if end <= 0 {
            return nil
        }
    }
    if end > l {
        end = l
    }
    return (*a)[begin:end]
}
