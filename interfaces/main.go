package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	Note "example.com/section6/note"
	Todo "example.com/section6/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputter interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	note, err := Note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	todoText := getUserInput("Enter Todo: ")
	todo, err := Todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}
	displayData(note)
	displayData(todo)
	err = saveData(note)
	if err != nil {
		return
	}
	saveData(todo)
}

func displayData(data displayer) {
	data.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Enter Title: ")
	content := getUserInput("Enter Note: ")
	return title, content
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

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("It's an integer", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("It's a float", floatVal)
		return
	}
	switch value.(type) {
	case string:
		fmt.Println("It's a string", value)
	case int:
		fmt.Println("It's an integer", value)
	default:
		fmt.Println("Unknown type", value)
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}

	fmt.Println("Saved successfully.")
	return nil
}
