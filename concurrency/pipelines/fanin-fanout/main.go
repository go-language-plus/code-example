package main

import (
	"fmt"
	"sync"
)

// 扇入和扇出
func main() {
	in := gen(2, 3)

	// 扇出：通常情况下多个函数可以同时从一个 channel 接收数据，直到 channel 关闭，这种情况被称作扇出。
	// 这是一种将工作分布给一组工作者的方法，目的是并行使用 CPU 和 I/O。
	// sq 里两个 goroutines 都从 in 中读取值
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
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

// 扇入：如果一个函数同时接收并处理多个 channel 输入并转化为一个输出 channel，直到所有的输入 channel 都关闭后，关闭输出 channel，这种情况就被称作扇入。
// 合并 channel 里的数据统一由一个 channel 传输
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 定义一个方法；传入一个通道；从该通道中取出值，发送到外部定义的 out 通道中
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done() // 完成后调用 wg.Done()
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

	// 注意哦，代码是会直接走到这而不会阻塞的；因为我们另起 goroutine 执行 output；而且另起了一个 goroutine 执行 wg.Wait() 而不是在当前 goroutine 做这些事
	// 不过这里并不会有问题，这里不是 main 函数；当前函数退出不会直接结束，因为
	// 我们在外面使用 for range 进行循环，for range 可直接在 channels 上面迭代，将依次从 channel 接收数据，当 channel 被关闭并且没有值可接收时会跳出循环。
	return out
}
