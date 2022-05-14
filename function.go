package main

func main() {
	sayHello()
	sayHelloTo("Stiven", "Katuuk")
	response := getHello("Stipen")

	println(response)

	firstName, lastName := getFullName()
	firstName2, lastName2 := getFullName2()

	println(firstName, lastName)
	println(firstName2, lastName2)
}

// non parameter
func sayHello() {
	println("Hello")
}

// with parameter
func sayHelloTo(firstName string, lastName string) {
	println("Hello", firstName, lastName)
}

// With Return
func getHello(name string) string {
	return "Hello " + name
}

// with multiple
func getFullName() (string, string) {
	return "Eko", "Stiven"
}

// Named Return
func getFullName2() (firstName string, lastName string) {
	firstName = "Eko"
	lastName = "Stiven"

	return
}
