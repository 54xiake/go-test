package main

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"time"
)

func main() {
	tk := toolbox.NewTask("myTask", "0/10 * * * * *", func() error { fmt.Println("hello world"); return nil })
	//err := tk.Run()
	//if err != nil {
	//	fmt.Println(err)
	//}
	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()
	time.Sleep(60 * time.Second)
	toolbox.StopTask()
}
