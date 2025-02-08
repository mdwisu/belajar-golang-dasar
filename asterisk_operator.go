package main

import "fmt"
type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}