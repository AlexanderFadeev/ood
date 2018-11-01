package main

import (
	"fmt"
	"os"

	"ood/lab5/editor"
	"ood/lab5/menu"

	"github.com/sirupsen/logrus"
)

func main() {
	e, err := editor.New()
	checkError(err)
	defer e.Release()

	m := menu.New(os.Stdin, os.Stdout)

	m.AddCommand("InsertParagraph", "Insert paragraph; args: <pos>|end <text>", e.InsertParagraph)
	m.AddCommand("InsertImage", "Insert image; args: <pos>|end <width> <height> <path>", e.InsertImage)
	m.AddCommand("SetTitle", "Set document title; args: <title>", e.SetTitle)
	m.AddCommand("List", "List elements from document", e.List)
	m.AddCommand("ReplaceText", "Replace text of a paragraph; args: <pos> <text>", e.ReplaceText)
	m.AddCommand("ResizeImage", "Resize an image; args: <pos> <width> <height>", e.ResizeImage)
	m.AddCommand("DeleteItem", "Delete an element; args: <pos>", e.DeleteItem)
	m.AddCommandWithoutArgs("Undo", "Undo command", e.Undo)
	m.AddCommandWithoutArgs("Redo", "Redo undone command", e.Redo)
	m.AddCommand("Save", "Save document as HTML; args: <path>", e.Save)

	m.AddCommandWithoutArgs("Exit", "Exit", func() error { m.Exit(); return nil })
	m.AddCommandWithoutArgs("Help", "Show help", func() error { m.Help(); return nil })

	m.SetDefaultHandler(unknownCommand)

	m.Run()
}

func checkError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

func unknownCommand() error {
	fmt.Println("Unknown command; Use `Help` command to show commands list")
	return nil
}
