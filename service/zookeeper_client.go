package main

import (
	"context"
	"flag"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

var (
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx", "prefix path")
)

func main() {
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	option.Group = "test"

	d := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkAddr}, nil)
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

		err1 := xclient.Broadcast(context.Background(), "Add", args, reply)
		if err1 != nil {
			log.Printf("ERROR failed to call: %v", err)
		}
		log.Printf("%d + %d = %d", args.A, args.B, reply.C)

		time.Sleep(1e9)
	}
}
