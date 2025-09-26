package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const DB_FILE = "balance.json"

func readFromFile() (float64, error) {
	if _, err := os.Stat(DB_FILE); err != nil {
		writeToFile(0)
	}
	data, err := os.ReadFile(DB_FILE)
	if err != nil {
		return 0, errors.New("ERROR: could not read balance file")
	}
	var result map[string]float64
	err = json.Unmarshal(data, &result)
	if err != nil {
		return 0, errors.New("ERROR: could not parse balance file")
	}
	balance := result["balance"]
	return balance, nil
}

func writeToFile(balance float64) error {
	data, err := json.Marshal(map[string]float64{"balance": balance})
	if err != nil {
		return errors.New("ERROR: could not encode balance data")
	}
	err = os.WriteFile(DB_FILE, data, 0644)
	if err != nil {
		return errors.New("ERROR: could not write balance file")
	}
	return nil
}

func main() {
	fmt.Println("Welcome to the Bank Application!")
	balance, err := readFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Your current balance is: $%.2f\n", balance)

	var choice int
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 3 {
			fmt.Println("Thank you for using the Bank Application. Goodbye!")
			break
		}

		var amount float64
		fmt.Print("Enter the amount: ")
		fmt.Scanln(&amount)

		switch choice {
		case 1:
			balance += amount
			fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, balance)
		case 2:
			if amount > balance {
				fmt.Println("ERROR: Insufficient funds for withdrawal.")
			} else {
				balance -= amount
				fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, balance)
			}
		default:
			fmt.Println("ERROR: Invalid choice. Please try again.")
			continue
		}

		err = writeToFile(balance)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
