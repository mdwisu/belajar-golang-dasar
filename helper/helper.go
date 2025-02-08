package helper

import "fmt"

var version = "1.0.0" // tidak bisa di akses di luar harus hurup besar
var Application = "golang"

func sayGoodBye(name string) string { // tidak bisa di akses di luar harus hurup besar
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Dwi")
	fmt.Println(version)
}