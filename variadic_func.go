package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	// total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := sumAll(number...)
	fmt.Println(total)
}