package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	println("Title:", note.Title)
	println("Content:", note.Content)
	println("Created At:", note.CreatedAt.Format("2006-01-02 15:04:05"))
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_"))

	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName+".json", jsonData, 0644)
}
