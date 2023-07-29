package main

import "fmt"

func add(num1, num2 int) int {
	res := num1 + num2
	return res
}

func main() {
	fmt.Println("First Number")
	var num1 int
	_, err := fmt.Scan(&num1)

	if err != nil {
		fmt.Println("Error reading number", err)
		return
	}

	fmt.Println("Second Number")
	var num2 int
	_, err2 := fmt.Scan(&num2)

	if err2 != nil {
		fmt.Println("Error reading number", err2)
		return
	}

	sum_of_two := add(num1, num2)

	print(sum_of_two)

}
