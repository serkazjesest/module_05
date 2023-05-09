package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	n1, n2 := 44, 45
	go fmt.Printf("\rFibonacci(%d) = %d\n", n1, fib(n1))
	fmt.Printf("\rFibonacci(%d) = %d\n", n2, fib(n2))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
