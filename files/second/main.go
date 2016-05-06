package main

import "fmt"

func main() {
	fmt.Printf("(P1)Hello")

	go World()

	var input string
	fmt.Scanln(&input)
}

func World() {
	fmt.Printf(" world(P2)\n")
}
