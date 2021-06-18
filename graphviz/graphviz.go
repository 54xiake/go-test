package main

import (
	"bytes"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	graphAst, _ := gographviz.Parse([]byte(`digraph G{}`))
	graph := gographviz.NewGraph()
	gographviz.Analyse(graphAst, graph)
	graph.AddNode("G", "a", map[string]string{"color": "red"})
	graph.AddNode("G", "b", nil)
	graph.AddNode("G", "c", nil)
	graph.AddEdge("a", "b", true, nil)
	graph.AddEdge("a", "c", true, map[string]string{"arrowhead": "diamond"})
	fmt.Println(graph.String())

	str, _ := os.Getwd()
	fmt.Println(str)
	// 输出文件
	ioutil.WriteFile("11.gv", []byte(graph.String()), 0666)

	// 产生图片
	system("dot 11.gv -T svg -o 12.svg")
}

//调用系统指令的方法，参数s 就是调用的shell命令
func system(s string) {
	cmd := exec.Command(`/bin/sh`, `-c`, s) //调用Command函数
	var out bytes.Buffer                    //缓冲字节

	cmd.Stdout = &out //标准输出
	err := cmd.Run()  //运行指令 ，做判断
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String()) //输出执行结果
}
