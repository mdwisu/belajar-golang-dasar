package main

import "fmt"

func main() {
	a := 20
	b := 20
	c := a + b
	d := a - b
	e := a * b
	f := a / b
	
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// augmented assignment
	i := 10
	i+= 10 // i = i + 10
	fmt.Println(i)
	i-= 10 // i = i - 10
	fmt.Println(i)
	i*= 10 // i = i * 10
	fmt.Println(i)
	i/= 10 // i = i / 10
	fmt.Println(i)

	// unary operator
	j := 10
	j++ // j = j + 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}