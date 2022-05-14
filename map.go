package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John",
		"address": "123 Main St",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
