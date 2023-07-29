package main

import "fmt"

func reverseNumber(num int) int {
	reversed := 0
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return reversed
}

func main() {
	fmt.Println("Enter the number")
	var num int
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Error reading number:", err)
		return
	}

	reversedNum := reverseNumber(num)
	fmt.Println("Reversed number:", reversedNum)
}
