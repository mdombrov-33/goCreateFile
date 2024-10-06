package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t Todo) Display() {
	fmt.Println("Content:", t.Text)
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("input cannot be empty")
	}

	return Todo{
		Text: content,
	}, nil
}

func (n Todo) SaveToFile() error {
	fileName := "todo.json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}
