package main

import "fmt"

var i = 1

func main() {
	fact := factorial(5)
	fmt.Println("Factorial is:", fact)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
