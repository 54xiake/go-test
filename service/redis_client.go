package main

import (
	"context"
	"flag"
	"log"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var (
	redisAddr = flag.String("redisAddr", "localhost:6379", "redis address")
	basePath  = flag.String("base", "/rpcx", "prefix path")
)

func main() {
	flag.Parse()

	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second

	d := client.NewRedisDiscovery(*basePath, "Arith", []string{*redisAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}
		//err := xclient.Call(context.Background(), "Mul", args, reply)
		//err := xclient.Fork(context.Background(), "Mul", args, reply)
		err := xclient.Broadcast(context.Background(), "Mul", args, reply)

		if err != nil {
			log.Printf("ERROR failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}

}
