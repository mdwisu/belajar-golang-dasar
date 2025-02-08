package main

import "fmt"

type hasName interface {
	getName() string
}
func SayHello(value hasName) {
	fmt.Println("Hello", value.getName())
}
type Person struct {
	Name string
}
func (person Person) getName() string {
	return person.Name
}
type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}
func main() {
	person := Person{Name: "Muhammad Dwi Susanto"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}