package main

import (
	"fmt"
	"sync"
)

// var a []int

// 需要指定通道中的元素的类型
var b chan int
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	b = make(chan int) // 不带缓冲区通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道中取到了", x)
	}()
	b <- 10 // hang住了
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)
	b = make(chan int, 16) // 带缓冲区通道初始化
	b <- 10
	b <- 20
	fmt.Println("10发送到通道b中了...")
	x := <-b
	fmt.Println("后台goroutine从通道中取到了", x)
	x = <-b
	fmt.Println("后台goroutine从通道中取到了", x)
	close(b)
}

func main() {
	noBufChannel()
	bufChannel()

	// 已经关闭的通道还是能取到值的
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	for x := range ch1 {
		fmt.Println(x)
	}
	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok)

}
