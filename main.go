package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mdombrov-33/gocreatefile/note"
)

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note body:")

	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()

	err = userNote.SaveToFile()

	if err != nil {
		fmt.Println("Saving the node failed")
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
