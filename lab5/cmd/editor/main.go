package main

import (
	"fmt"
	"os"

	"ood/lab5/editor"
	"ood/lab5/editor_handlers"
	"ood/lab5/menu"

	"github.com/sirupsen/logrus"
)

func checkError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	e, err := editor.New()
	checkError(err)

	m := menu.New(os.Stdin, os.Stdout)
	handlers := editor_handlers.New(e)

	m.AddCommand("InsertParagraph", "Insert paragraph; args: <pos>|end <text>", handlers.InsertParagraph)
	m.AddCommand("InsertImage", "Insert image; args: <pos>|end <width> <height> <path>", handlers.InsertImage)
	m.AddCommand("SetTitle", "Set document title; args: <title>", handlers.SetTitle)
	m.AddCommand("List", "List elements from document", handlers.List)
	m.AddCommand("ReplaceText", "Replace text of a paragraph; args: <pos> <text>", handlers.ReplaceText)
	m.AddCommand("ResizeImage", "Resize an image; args: <pos> <width> <height>", handlers.ResizeImage)
	m.AddCommand("DeleteItem", "Delete an element; args: <pos>", handlers.DeleteItem)
	m.AddCommandWithoutArgs("Undo", "Undo command", handlers.Undo)
	m.AddCommandWithoutArgs("Redo", "Redo undone command", handlers.Redo)
	m.AddCommand("Save", "Save document as HTML; args: <path>", handlers.Save)

	m.AddCommandWithoutArgs("Exit", "Exit", func() error { m.Exit(); return nil })
	m.AddCommandWithoutArgs("Help", "Show help", func() error { m.Help(); return nil })

	m.SetDefaultHandler(UnknownCommand)

	m.Run()
}

func UnknownCommand() error {
	fmt.Println("Unknown command; Use `Help` command to show commands list")
	return nil
}
