package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	slice := []string{"Eko", "Kurniawan", "Khanedi"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println(index, name)
	}
}
