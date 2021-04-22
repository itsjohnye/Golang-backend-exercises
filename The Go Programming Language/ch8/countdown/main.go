package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launched!")
}

func main() {
	fmt.Println("Commecing countdown.")
	tick := time.Tick(1 * time.Second) //time.Tick函数返回一个通道

	abort := make(chan struct{})

	//go func()作为监控者
	go func() {
		os.Stdin.Read(make([]byte, 1)) //读取单个字节
		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			//do not thing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}

	}
	launch()
}
