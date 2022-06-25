package main

func main() {
	var counter int8 = 0
	increment := func() {
		println("Increment")
		counter++
	}

	increment()
	increment()

	println(counter);
}