package main

func sayhello(firstname string, lastname string) string {
	println("Hello " + firstname + " " + lastname)
	return "Hello " + firstname + " " + lastname
}

// multiple return value
func getFullName() (string, string) {
	return "muhammad", "dwi"
}

func main() {
	returnvalue := sayhello("Muhammad", "Dwi Susanto")
	println(returnvalue)

	firstname, lastname := getFullName()
	println(firstname, lastname)
	// ignore return value
	_, _ = getFullName()
}