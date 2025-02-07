package main

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		println("You are blocked", name)
	} else {
		println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("dwi", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}