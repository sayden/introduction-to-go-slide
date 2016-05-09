package main

import "fmt"

func main() {
	for i := 0; i < 10000; i++ {
		go func(j int) {
			fmt.Printf("Proccess #%d\n", j)
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
