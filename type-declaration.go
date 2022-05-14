package main

import "fmt"

func main() {
	type noKTP string

	var noKTP1 noKTP = "123456789"
	fmt.Println(noKTP1)
}
