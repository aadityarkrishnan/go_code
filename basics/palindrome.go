package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	println("Enter the string")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if isPalindrome(input) {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}

}
