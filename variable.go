package main

import "fmt"

func main() {
	name := "stiven"

	name = "Stiven"
	fmt.Println("Nama saya adalah ", name)

	var (
		umur   = 18
		alamat = "Jakarta"
	)

	fmt.Println("Umur saya adalah ", umur)
	fmt.Println("Alamat saya adalah ", alamat)
}
