package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var dwi Customer
	dwi.Name = "Muhammad Dwi Susanto"
	dwi.Address = "Jl. Raya No. 1"
	dwi.Age = 20
	fmt.Println(dwi)

	// struct literal
	joko := Customer{
		Name:    "Joko",
		Address: "Jl. Raya No. 2",
		Age:     30,
	}
	fmt.Println(joko)
	budi := Customer{"budi", "Jl. Raya No. 3", 40}
	fmt.Println(budi)

	budi.sayHello("Dwi")
}