/*
   fileName: goQueue
   author: diogoxiang@qq.com
   date: 2019/7/7
*/
package main

import (
	"fmt"
	"github.com/nickalie/go-queue"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go consumer(i + 1)
	}

	producer()
}

func producer() {
	i := 0
	for {
		i++
		queue.Put("messages", fmt.Sprintf("message %d", i))
		time.Sleep(time.Second)
	}
}

func consumer(index int) {
	for {
		var message string
		queue.Get("messages", &message)

		fmt.Printf("Consumer %d got a message: %s\n", index, message)
		time.Sleep(2 * time.Second)
	}
}
