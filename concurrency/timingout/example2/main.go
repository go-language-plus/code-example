package main

import (
	"fmt"
	"time"
)

var ch chan int

// Example 2，超时控制
func main() {
	select {
	case m := <-ch:
		// a read from ch has occurred
		doSomething(m)
	case <-time.After(10 * time.Second):
		// the read from ch has timed out
		fmt.Println("timed out")
	}
}

func doSomething(int) {}
