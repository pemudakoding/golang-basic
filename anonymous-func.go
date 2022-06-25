package main

import (
	"fmt"
)

func main() {
	blacklist := func(name string) bool {
		return name == "Stiven"
	}

	register("Stiven", blacklist)
	register("Eko", blacklist)
}

type Blacklist  func(string) bool

func register(name string, filter Blacklist) {
	if filter(name) {
		fmt.Println("You're blocked", name)
	}else {
		fmt.Println("Welcome", name)
	}
}