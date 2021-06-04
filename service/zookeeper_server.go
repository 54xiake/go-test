package main

import (
	"flag"
	graphite "github.com/cyberdelia/go-metrics-graphite"
	"github.com/rcrowley/go-metrics"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"net"
	"time"
)

var (
	addr     = flag.String("addr", "localhost:8972", "client address")
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	p := serverplugin.NewMetricsPlugin(metrics.DefaultRegistry)
	s.Plugins.Add(p)
	startMetrics()

	addRegistryPlugin(s)

	//s.RegisterName("Arith", new(example.Arith), "group=test")
	s.RegisterName("Arith", new(example.Arith), "group=test&state=active&tps=0")
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}

func startMetrics() {
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Second)

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:2003")
	go graphite.Graphite(metrics.DefaultRegistry, time.Second, "rpcx.services.host.127_0_0_1", addr)

}
