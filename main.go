package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mdombrov-33/gocreatefile/note"
	"github.com/mdombrov-33/gocreatefile/todo"
)

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note body:")

	return title, content
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

	err = todo.SaveToFile()

	if err != nil {
		fmt.Println("Saving the todo failed")
		return
	}

	fmt.Println("Saving successful!")

	userNote.Display()

	err = userNote.SaveToFile()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving successful!")
}

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
