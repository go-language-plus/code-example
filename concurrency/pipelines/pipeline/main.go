package main

import "fmt"

// Example of pipelines
func main() {
	// 建立管道
	c := gen(2, 3)
	out := sq(c)

	// 消费
	// 我们传进去两个值，把这两个值通过管道后的结果取出（从 sq 返回的通道中取）
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9

	// 我们甚至可以嵌套着这么写
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 会输出 16 和 81
	}
}

// 将传进来的 nums 扔进 out 通道中，返回这个通道
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 从传进来的 in 通道取值，将值做一次数学计算之后扔到一个 out 通道中，返回这个 out 通道
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
