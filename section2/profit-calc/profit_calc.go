package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getUserInput(msg string) float64 {
	fmt.Print(msg)
	var input float64
	fmt.Scan(&input)
	if input <= 0 {
		fmt.Println("Input must be a positive number.")
		return getUserInput(msg)
	}
	return input
}

func calcProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = profit / revenue
	return ebt, profit, ratio
}

func saveToFile(ebt, profit, ratio float64) error {
	contents := map[string]float64{
		"ebt":    ebt,
		"profit": profit,
		"ratio":  ratio,
	}
	jsonData, err := json.Marshal(contents)
	if err != nil {
		return err
	}
	if os.WriteFile("profit.json", jsonData, 0644) != nil {
		return err
	}
	return nil
}

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calcProfit(revenue, expenses, taxRate)
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	if err := saveToFile(ebt, profit, ratio); err != nil {
		panic(err)
	}
}
