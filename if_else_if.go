// if expresion

package main

import "fmt"

func main()  {
	var name string = "Muhammad Dwi Susanto"

	if name == "Muhammad Dwi Susanto" {
		fmt.Println("Hello World")
	} else if name == "Muhammad Dwi" {
		fmt.Println("Hello World 1")
	} else {
		fmt.Println("Hello World 2")
	}

	if length := len(name); length > 30 {
		fmt.Println("Nama Terlalu Panjang")
	} else if length > 15 {
		fmt.Println("Nama Sudah Benar benar")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}