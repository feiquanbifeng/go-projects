// Read json data from json file
// write json data to file

package json

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "path/filepath"
)

// Define a Quieter interface
type Quieter interface {
    Get(k string) interface{}
    Set(k string, v interface{})
    Len() int
    Load(fn func()) error
    Save(fn func()) error
}

// Define a quiet struct
type quiet struct {
    filename string
    data     map[string]interface{}
}

// Instance of quiet
func NewQuiet(filename string) (Quieter, error) {
    wd, err := os.Getwd()
    if err != nil {
        return nil, err
    }
    abs := filepath.Join(wd, filename)
    return &quiet{
        filename: abs,
        data:     make(map[string]interface{}),
    }, nil
}

// Get value by key
func (q *quiet) Get(k string) interface{} {
    return q.data[k]
}

// Set key value pair
func (q *quiet) Set(k string, v interface{}) {
    if q.data == nil {
        q.data = make(map[string]interface{})
    }
    q.data[k] = v
}

// Return data length
func (q *quiet) Len() int {
    return len(q.data)
}

// Load data
func (q *quiet) Load(fn func()) error {
    byt, err := ioutil.ReadFile(q.filename)
    if err != nil {
        return err
    }
    if err := json.Unmarshal(byt, &q.data); err != nil {
        return err
    }
    if fn != nil {
        fn()
    }
    return nil
}

// Save data
func (q *quiet) Save(fn func()) error {
    byt, err := json.Marshal(q.data)
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(q.filename, byt, 0655)
    if err != nil {
        return err
    }
    if fn != nil {
        fn()
    }
    return nil
}
