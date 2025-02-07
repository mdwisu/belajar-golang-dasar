package main

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	println("hello", filteredName)
}
func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("anjing", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("dwi", filter)
}
