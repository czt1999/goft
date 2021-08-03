package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 创建通道
	var ch = make(chan string, 10)
	for i := 0; i < 4; i++ {
		go producer(i, ch)
		go consumer(i, ch)
	}
	go func(ch chan string) {
		for {
			time.Sleep(time.Second * (time.Duration(3)))
			fmt.Printf(">>> len(ch): %d \n", len(ch))
		}
	}(ch)

	time.Sleep(time.Minute)
}

//
// 生产者函数
//
func producer(id int, ch chan string) {
	rand.Seed(int64(id))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * (time.Duration(rand.Intn(6))))
		s := fmt.Sprintf("[%d]:%d", id, i)
		ch <- s
	}
}

//
// 消费者函数
//
func consumer(id int, ch chan string) {
	for s := range ch {
		fmt.Printf("Consumer [%d] gets message:{ %s } \n", id, s)
		time.Sleep(time.Second * (time.Duration(rand.Intn(10))))
	}
}
