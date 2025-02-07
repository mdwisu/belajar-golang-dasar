// function value
package main

func getGoodBye(name string) string {
	if name == "" {
		return "Good Bye"
	} else {
		return "Good Bye " + name
	}
}

func main() {
	goodBye := getGoodBye
	println(goodBye("Muhammad Dwi Susanto"))
}
