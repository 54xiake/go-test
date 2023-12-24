package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	go HandelRequest(ctx)
//	time.Sleep(5 * time.Second)
//	fmt.Println("It's time to stop all sub goroutines!")
//	cancel()
//	//Just for test whether sub goroutines exit or not
//	time.Sleep(5 * time.Second)
//}

func main() {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	go func() { // run the work in the background
		if err := work(ctx, "D:\\myweb\\go-test\\basic\\example.txt"); err != nil {
			log.Println(err)
		}
	}()

	// perform some operation and that causes error
	time.Sleep(time.Millisecond * 150)
	if true { // err != nil
		stop()
	}
	time.Sleep(time.Second)
}

func work(ctx context.Context, filename string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		time.Sleep(time.Millisecond * 100)
		log.Print(line) // do something with the line
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
	return nil
}
