package main

import (
	"fmt"
	"sync"
)

// 基于 Pipelines 扇入扇出的代码；我们做些修改
func main() {
	// 创建一个 done channel
	// pipeline 退出时会 close done channel；通知所有 goroutines 退出
	done := make(chan struct{})
	defer close(done)

	in := gen(2, 3)

	c1 := sq(in)
	c2 := sq(in)

	out := merge(done, c1, c2)
	fmt.Println(<-out) // 取一个，我就退出了；
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

// 第一个参数为 done channel；第二个参数开始才是要合并的 channel
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 往 out 中发送值
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs)) // 要合并的（传进来的）通道数

	// 每个传进来的通道都调用 output 方法；即所有通道都通过 output 方法往 out 通道发送数据；最终合并为 out 通道返回
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait() // 等到所有调用完成后，关闭 out（ch 关闭后还可以取值；只是不能发送值）
		close(out)
	}()

	return out
}
