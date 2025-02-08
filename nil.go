package main

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{"name": name}
}
func main() {
	data := NewMap("")
	if data == nil {
		println("Data Kosong")
	} else {
		println(data["name"])
	}
}