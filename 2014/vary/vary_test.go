package vary

import (
    "net/http"
    // "io/ioutil"
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
