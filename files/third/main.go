package main

import "fmt"

func main() {
	for i := 0; i < 10000; i++ {
		go func(iter int) {
			fmt.Printf("Proccess #%d\n", iter)
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
