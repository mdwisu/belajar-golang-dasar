package main

import "fmt"
type Address struct {
	City, Province, Country string
}

func main() {
	alamat1 := new(Address)
	alamat1.City = "Bandung"
	alamat1.Province = "Jawa Barat"
	alamat1.Country = "Indonesia"
	fmt.Println(alamat1)
}