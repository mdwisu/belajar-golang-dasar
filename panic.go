package main

import "fmt"

func endApp() {
	println("Aplikasi selesai")
	message := recover()
	fmt.Println("Error : ", message)
}
func runningApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}

	println("Aplikasi berjalan")
}
func main() {
	runningApp(true)
	println("dwi")
}