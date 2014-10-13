// Redis operator
// use pool

package saver

import (
    "github.com/garyburd/redigo/redis"

    "sync"
    "time"
)

var (
    pool          *redis.Pool
    redisServer   = "pub-redis-10230.us-east-1-4.2.ec2.garantiadata.com:10230"
    redisPassword = "realtime"
    prefix        = "db_"
)

func init() {
    pool = newPool(redisServer, redisPassword)
}

// return a pool with maxIdle is 3 and validate with password
func newPool(server, password string) *redis.Pool {
    return &redis.Pool{
        MaxIdle:     3,
        IdleTimeout: 240 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            if _, err := c.Do("AUTH", password); err != nil {
                c.Close()
                return nil, err
            }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
}

// Define command struct
// maybe distribute environment
type command struct {
    sync.RWMutex
    conn redis.Conn
}

// New a command
func NewCommand() *command {
    return &command{
        conn: pool.Get(),
    }
}

// Store different product user info
func (c *command) zadd(i int, member string) {
    c.conn.Send("ZADD", prefix+"users", i, member)
}

// Store offline message
func (c *command) list(member string) {
    c.conn.Send("LPUSH", prefix+"slots", member)
}

// Store message entity key-value
// key is id and value is content
func (c *command) dict(i, member string) {
    c.conn.Send("HSET", prefix+"buckets", i, member)
}

// Close the conn release resource
func (c *command) close() error {
    return c.conn.Close()
}
