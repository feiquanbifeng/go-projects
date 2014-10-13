// RPC operator
// listen port 1234

package saver

import (
    "log"
    "net"
    "net/http"
    "net/rpc"
)

const port = ":1234"

func init() {
    server := new(Server)
    rpc.Register(server)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", port)
    if e != nil {
        log.Fatal("listen error:", e)
    }
    go http.Serve(l, nil)
}

type Message struct {
    Id  int
    Msg string
}

type Server struct{}

func (s *Server) AddMessage(args *Message, result *string) error {
    c := NewCommand()
    c.zadd(args.Id, args.Msg)
    c.close()
    return nil
}
