package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

func main() {
	log.Println("Redis:")
	conn, err := redisConn("127.0.0.1:6379", "", "1")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Println(conn.dbid)
	test(conn)
}

func test(conn *RedisConn) {
	//conn.Do("SET", "xxx", 1)
	//conn.Do("GET", "xxx")
	if xxx, err := redis.Int(conn.Do("INCR", "xxx")); err == nil {
		log.Println("xxx:", xxx)
	}
	conn.Close()
	//conn.FlushClose()
}

////////////////////////////////////////////////////////////////
type RedisConn struct {
	dbid string
	redis.Conn
}

func (r *RedisConn) FlushClose() error {
	if r.dbid != "" {
		if _, err := r.Conn.Do("SELECT", r.dbid); err != nil {
			return nil
		}
	}
	if _, err := r.Conn.Do("FLUSHDB"); err != nil {
		return err
	}
	return r.Conn.Close()
}

func (r *RedisConn) Close() error {
	return r.Conn.Close()
}

func redisConn(host, password, db string) (*RedisConn, error) {
	if host == "" {
		host = ":6379"
	}

	dialReadTimeout := redis.DialReadTimeout(1 * time.Second)
	dialWriteTimeout := redis.DialWriteTimeout(1 * time.Second)

	conn, err := redis.Dial("tcp", host, dialReadTimeout, dialWriteTimeout)
	//conn, err := redis.DialTimeout("tcp", host, 0, 1*time.Second, 1*time.Second)
	if err != nil {
		return nil, err
	}

	if password != "" {
		if _, err := conn.Do("AUTH", password); err != nil {
			conn.Close()
			return nil, err
		}
	}

	if db != "" {
		if _, err := conn.Do("SELECT", db); err != nil {
			conn.Close()
			return nil, err
		}
	}

	return &RedisConn{dbid: db, Conn: conn}, nil
}
