package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mdombrov-33/gocreatefile/note"
	"github.com/mdombrov-33/gocreatefile/todo"
)

type saver interface {
	SaveToFile() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	todo.Display()

	err = saveToFile(todo)

	if err != nil {
		return
	}

	userNote.Display()

	err = saveToFile(userNote)

	if err != nil {
		return
	}

}

func saveToFile(data saver) error {
	err := data.SaveToFile()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving successful!")
	return nil
}

// get input
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	bufioReader := bufio.NewReader(os.Stdin)
	text, err := bufioReader.ReadString('\n')

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// get note data
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note body:")

	return title, content
}
