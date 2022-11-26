package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)
	}
}
func main() {
	done := make(chan bool, 1)

	go func() {
		slow("hello")
		done <- true
	}()
	slow("sing")
	<-done
	fmt.Println("all task done")
}
