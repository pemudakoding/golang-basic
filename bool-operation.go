package main

func main() {
	var a, b bool

	a = true
	b = false

	if a && b {
		println("a && b")
	}
	if a || b {
		println("a || b")
	}
	if !a {
		println("!a")
	}
}
