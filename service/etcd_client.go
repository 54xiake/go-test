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
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx", "prefix path")
)

func main() {
	d := client.NewEtcdDiscovery(*basePath, "Arith", []string{*etcdAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
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
