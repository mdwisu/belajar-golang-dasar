// function array

package main

import "fmt"

func main()  {
	var names [3]string

	names[0] = "Muhammad Dwi Susanto"
	names[1] = "Muhammad Dwi Susanto"
	names[2] = "Muhammad Dwi Susanto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names))

	var values = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(values[0])
	fmt.Println(len(values))
}