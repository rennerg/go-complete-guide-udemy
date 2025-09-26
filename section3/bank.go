package main

import (
	"fmt"

	"example.com/m/utilities"
	"github.com/Pallinder/go-randomdata"
)

const DB_FILE = "balance.json"

func main() {
	fmt.Println("Welcome to the " + randomdata.SillyName() + " Bank!")
	fmt.Println("  Phone #:", randomdata.PhoneNumber())
	fmt.Println("  Address:", randomdata.Address())
	fmt.Println()
	balance, err := utilities.ReadFromFile(DB_FILE, "balance")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Your current balance is: $%.2f\n", balance)

	var choice int
	for {
		presentOptions()
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

		err = utilities.WriteToFile(DB_FILE, "balance", balance)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
