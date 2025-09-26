package main

import "fmt"

func main() {
	prices := []float64{19.99, 29.99, 4.99, 49.99}

	fmt.Println("Prices:", prices)
	prices = append(prices, 9.99)
	fmt.Println("Updated Prices:", prices)
	prices = prices[0 : len(prices)-1]
	fmt.Println("After Removing Last Price:", prices)
}
