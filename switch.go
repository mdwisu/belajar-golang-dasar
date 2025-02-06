// switch

package main

import "fmt"

func main()  {
	var name string = "asdf"

	switch name {
	case "Muhammad":
		fmt.Println("Hello Muhammad")
	case "Dwi":
		fmt.Println("Hello Dwi")
	case "Susanto":
		fmt.Println("Hello Susanto")
	default:
		fmt.Println("Hello World")
	}
	// short
	switch length := len(name); length > 10 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	default:
		fmt.Println("Nama Sudah Benar")
	}
	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}