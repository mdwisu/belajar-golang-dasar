// for

package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// for range
	names := []string{"Muhammad", "Dwi", "Susanto"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for index, name := range names {
		fmt.Println(index, name)
	}
	angka := []int{1, 2, 3, 4, 5}
	for _, v := range angka {
		fmt.Println(v)
	}
}