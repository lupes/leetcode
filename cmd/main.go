package cmd

import (
	"fmt"
)

func main() {
	fmt.Println(fib1(50))
	//for i := 1; i < 50; i++ {
	//	fmt.Println(fib2(int64(i)))
	//}
}

func fib1(n int32) int32 {
	if n > 2 {
		return fib1(n-1) + fib1(n-2)
	} else if n == 1 || n == 2 {
		return 1
	} else {
		panic("must bigger 0")
	}
}

func fib2(n int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	var (
		i int64 = 3
		a int64 = 1
		b int64 = 1
	)
	for ; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
