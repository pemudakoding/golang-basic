package main

func resolveForFactorial(value int) int {
	if value == 1 {
		return 1
	}

	return value * resolveForFactorial(value - 1)
}

func main() {
	println(resolveForFactorial(10));
}