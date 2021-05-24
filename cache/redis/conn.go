package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool        *redis.Pool
	redisHost   = "127.0.0.1:6379"
	redisPasswd = "entry_test_4453"
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTImeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) { // connect to redis server
			conn, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Printf("redis conn faled %s\n", err.Error())
				return nil, err
			}

			if _, err := conn.Do("AUTH", redisPasswd); err != nil {
				conn.Close()
				return nil, err
			}

			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error { // check redis connection normal or not
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

// Pool exposes redis.Pool to outside of package
func Pool() *redis.Pool {
	return pool
}
