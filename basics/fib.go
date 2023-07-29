package main

import (
	"fmt"
)

func generateFibonacci(limit int) []int {
	fib := []int{0, 1}
	for i := 2; ; i++ {
		nextFib := fib[i-1] + fib[i-2]
		if nextFib > limit {
			break
		}
		fib = append(fib, nextFib)
	}
	return fib
}

func main() {
	var limit int
	println("Enter the limit")
	_, err := fmt.Scan(&limit)
	if err != nil {
		println("Some error has created", err)
	}

	fibonacciSeries := generateFibonacci(limit)
	fmt.Println("Fibonacci series up to", limit, ":", fibonacciSeries)
}
