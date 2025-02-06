package main

import "fmt"

func main() {
	type NoKTP string // type NoKTP adalah alias untuk string

	var noKTP NoKTP = "1234567890"
	fmt.Println(noKTP)

	var contoh string = "19238172937192"
	var contoh2 NoKTP = NoKTP(contoh)
	fmt.Println(contoh2)
}