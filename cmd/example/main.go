package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("worker", workerId, "recebeu", res)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	canal := make(chan int)
	for i := 0; i < 2; i++ {
		go worker(i, canal)
	}

	for i := 0; i < 10; i++ {
		canal <- i
	}

	time.Sleep(time.Second * 20)
}
