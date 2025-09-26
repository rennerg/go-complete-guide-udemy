package main

import (
	"fmt"

	"example.com/section5/user"
)

func main() {
	firstName := getUserData("Enter First Name: ")
	lastName := getUserData("Enter Last Name: ")
	birthdate := getUserData("Enter Birthdate (YYYY-MM-DD): ")

	u, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.DisplayInfo()
	u.ClearUserName()
	u.DisplayInfo()

	a := user.NewAdmin("admin@example.com", "securepassword")
	a.DisplayInfo()
}

// Private methods

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scan(&input)
	return input
}
