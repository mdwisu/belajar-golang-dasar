package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// address2 := address1
	// address3 := &address1 //pointer

	// address2.City = "Jakarta"
	// address3.City = "Surabaya"

	// fmt.Println(address1)
	// fmt.Println(address2)
	// fmt.Println(address3)

	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}