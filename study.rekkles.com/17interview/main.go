package main

type a struct {
	val  int
	next *a
}

// 判断闭环
// x 一次走一步
// y 一次走两步
// 如果再某个时刻他们能相遇就有闭环

// n个台阶，一次走一步，一次走两步
// 有多少种走法
func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}

// n =1000时 如何优化
// 1 x
// 2 y
// 3 x+y
// 4
// 5
// 6
// 7
