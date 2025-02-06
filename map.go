// map

package main

import "fmt"

func main()  {
	// var person map[string]string = make(map[string]string)
	// var person map[string]string = map[string]string{}
	// person["name"] = "Muhammad Dwi Susanto"
	// person["age"] = "20"

	// fmt.Println(person["name"])
	// fmt.Println(person["age"])

	person:= map[string]string{
		"name": "Muhammad Dwi Susanto",
		"age": "20",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)

	// function map
	len := len(person)
	fmt.Println(len)
	person["address"] = "Jl. Raya No. 1"
	fmt.Println(person)
	delete(person, "name")
	fmt.Println(person)
}