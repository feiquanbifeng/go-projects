// bamboo utils
// since 2016-01-01 by jy

package utils

import (
    "bufio"
    "bytes"
    "io/ioutil"
    "strings"
)

const split string = "="

var configs map[string]*Config

// define the config struct
type Config struct {
    properties map[string]string
}

func (c *Config) GetValue(key string) string {
    return c.properties[key]
}

func NewConfig(filename string) *Config {
    if c, ok := configs[filename]; ok {
        return c
    }
    configs = make(map[string]*Config)
    conf := &Config{
        properties: readFile(filename),
    }
    configs[filename] = conf
    return conf
}

// Private method
func readFile(f string) map[string]string {

    c := make(map[string]string)
    buf, err := ioutil.ReadFile(f)
    if err != nil {
        panic(err)
    }
    s := bufio.NewScanner(bytes.NewBuffer(buf))
    for s.Scan() {
        text := s.Text()
        strs := strings.Split(text, split)
        key := strings.TrimSpace(strs[0])
        value := strings.TrimSpace(strs[1])
        c[key] = value
    }
    return c

}
