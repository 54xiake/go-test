package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")

	mu := sync.Mutex{}
	//go func() {
	//	for {
	//		fmt.Println("go")
	//	}
	//}()
	ch := make(chan int, 3)
	//defer close(ch)
	go func() {
		mu.Lock()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	select {
	case data, ok := <-ch:
		if ok {
			fmt.Println(data)
		}
		//mu.Unlock()
		//default:
		//	fmt.Println("default")

	}
}

func main() {

	//http.HandleFunc("/debug/pprof/", pprof.Index)
	//http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//http.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)
	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
