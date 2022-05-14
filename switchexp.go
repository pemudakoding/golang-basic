package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Stiven":
		fmt.Println("Hello Stiven")
	default:
		fmt.Println("Hello World")
	}
}
