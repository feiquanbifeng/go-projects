// read json data from json file
// write json data to file

package json

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "path/filepath"
)

// define a Quiet struct
type Quiet struct {
    filename string
    Data     map[string]interface{}
}

// instance of Quiet
func NewQuiet(filename string) (*Quiet, error) {
    wd, err := os.Getwd()
    if err != nil {
        return nil, err
    }
    abs := filepath.Join(wd, filename)
    return &Quiet{
        filename: abs,
        Data:     make(map[string]interface{}),
    }, nil
}

// load data
func (q *Quiet) Load(fn func()) error {
    byt, err := ioutil.ReadFile(q.filename)
    if err != nil {
        return err
    }
    if err := json.Unmarshal(byt, &q.Data); err != nil {
        return err
    }
    if fn != nil {
        fn()
    }
    return nil
}

// save data
func (q *Quiet) Save(fn func()) error {
    byt, err := json.Marshal(q.Data)
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
