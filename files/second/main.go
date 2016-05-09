package main

import "fmt"

func main() {
	fmt.Printf("(P1)Hello")

	go World()

}

func World() {
	fmt.Printf(" world(P2)\n")
}
