package vary

import (
    "net/http"
    "fmt"
    "testing"
)

func TestVary(t *testing.T) {
    req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
    client := &http.Client{
       // CheckRedirect: redirectPolicyFunc,
    }
    resp, _ := client.Do(req)
    Vary(resp, "User-Agent")
    // Append("Accept-Encoding", "Origin")

    defer resp.Body.Close()

    //body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(resp.Header["Vary"])
}

func TestAppend(t *testing.T) {
    a := Append("", "Origin", "User-Agent", "*", "Accept")
    if a == "*" {
        fmt.Println("ok")
    } else {
        fmt.Println("error")
    }
}
