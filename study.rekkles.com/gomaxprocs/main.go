package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	// M:N 把m个GOROUTINE分配给n个操作系统线程去执行
	// goroutine初始栈的大小是2k
	// 只有一个干活的
	// 默认CPU逻辑核心数，默认跑满整个CPU
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
