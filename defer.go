package main

func logging() {
	println("seleai")
}
func runApp() {
	defer logging() // akan di panggil di akhir
	println("run app")
}
func main() {
	runApp()
}