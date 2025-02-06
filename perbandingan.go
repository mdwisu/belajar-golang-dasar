// perbandingan

package main

import "fmt"

func main() {
	// perbandingan
	var name string = "Muhammad Dwi Susanto"
	var name2 string = "Muhammad Dwi Susanto"
	var result bool = name == name2
	fmt.Println(result)
	fmt.Println("1 == 2", 1 == 2) // false
	fmt.Println("1 != 2", 1 != 2) // true
	fmt.Println("1 > 2", 1 > 2)   // false
	fmt.Println("1 < 2", 1 < 2)   // true
	fmt.Println("1 >= 2", 1 >= 2) // false
	fmt.Println("1 <= 2", 1 <= 2) // true
}