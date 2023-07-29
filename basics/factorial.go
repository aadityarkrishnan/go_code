package main

import (
	"fmt"
)

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	var num int
	println("Enter the number")
	_, err := fmt.Scan(&num)

	if err != nil {
		println("Found some error", err)
		return
	}

	println("Factorial of", num, "is =", factorial(num))
}
