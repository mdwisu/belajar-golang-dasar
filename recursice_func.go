package main

func factorialLoop(v int) int {
	result := 1
	for i := v; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(v int) int {
	if v == 1 {
		return 1
	}
	return v * factorialRecursive(v-1)
}

func main() {
	println(factorialLoop(5))
	println(factorialRecursive(5))
}