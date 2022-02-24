package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutine是用户态线程
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
}

func f() {
	// 不加SEED每次运行结果会一样
	rand.Seed(time.Now().UnixNano())
	// 5577006791947779410 7
	// 6129484611666145821 9
	// 3916589616287113937 8
	// 605394647632969758 0
	// 894385949183117216 0
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // x>=0 x<10
		fmt.Println(0-r1, 0-r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

// 程序启动之后会创建一个住GOROUTINE去执行
func main() {
	// for i := 0; i < 100; i++ {
	// 	//开启一个单独的GOROUTINE去执行HELLO函数（任务）
	// 	// go hello(i)
	// 	go func() {
	// 		fmt.Println(i) // 变量查找
	// 	}()
	// }
	// fmt.Println("main")
	// time.Sleep(time.Second)
	// main 函数结束了 由main函数启动的GOROUTINE也结束了

	// f()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	// 等待wg的计数器减为0
	wg.Wait()

}
