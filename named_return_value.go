package main

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Dwi"
	lastName = "Susanto"
	return firstName, middleName, lastName
}
func main() {
	a, b, c := getCompleteName()

	println(a, b, c)
}
