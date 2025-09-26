package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	Text string `json:"content"`
}

func (todo Todo) Display() {
	println("Todo:", todo.Text)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content cannot be empty")
	}
	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Save() error {
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile("todo.json", jsonData, 0644)
}
