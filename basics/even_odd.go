package main

import "fmt"

func check_even_odd(num int) string {
	if num%2 == 0 {
		return "EVEN"
	} else {
		return "ODD"
	}
}

func main() {
	var num1 int
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Error in finding the type")
		return
	}
	println("Number ", num1, "is ", check_even_odd(num1))
}
