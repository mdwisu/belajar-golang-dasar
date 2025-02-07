package main

import "fmt"

// closures
func main(){
	counter := 0
	increment := func() int{
		counter++
		return counter
	}
	// jangan terlalu banyak menggunakan ini karena akan membingungkan, karena di pertengahan mengubah nilai counter
	fmt.Println(increment())
	fmt.Println(increment())
}