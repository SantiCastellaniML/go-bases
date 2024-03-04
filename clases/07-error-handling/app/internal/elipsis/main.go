package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}

	var result = Addition(sl...)

	fmt.Println(result)
}

func Addition(values ...int) int {
	var result int
	for _, value := range values {
		result += value
	}

	return result
}
