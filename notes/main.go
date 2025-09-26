package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	Note "example.com/notes/note"
)

func main() {
	title := getUserInput("Enter Title: ")
	content := getUserInput("Enter Note: ")
	note, err := Note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	note.Display()
	err = note.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	text := strings.TrimSuffix(input, "\n")
	text = strings.TrimSuffix(text, "\r") // For Windows compatibility

	return text
}
