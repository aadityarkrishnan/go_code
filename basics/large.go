package main

import "fmt"

func main() {
	fmt.Println("Largest of number")
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(findLargestNumber(array))
}

func findLargestNumber(arr []int) int {
	largest := arr[0]
	for _, num := range arr {
		if num > largest {
			largest = num
		}
	}
	return large
}
