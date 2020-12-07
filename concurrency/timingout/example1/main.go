package main

import "time"

var ch chan int

// Example 1，演示最基本的情况，实际使用时使用 example 2（time 包）
func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		// a read from ch has occurred
	case <-timeout:
		// the read from ch has timed out
	}
}
