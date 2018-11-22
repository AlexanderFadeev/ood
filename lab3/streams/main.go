package main

import (
	"io"
	"ood/lab3/streams/parser"
	"os"
)

func main() {
	//defer recoverPanic()

	p := parser.New()
	args, err := p.Parse(os.Args[1:])
	check(err)

	input, err := os.Open(args.InputFileName)
	check(err)
	inputDecorated := args.RDecorator.DecorateReader(input)
	defer inputDecorated.Close()

	output, err := os.OpenFile(args.OutputFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	check(err)
	outputDecorated := args.WDecorator.DecorateWriter(output)
	defer outputDecorated.Close()

	_, err = io.Copy(outputDecorated, inputDecorated)
	check(err)
}

func check(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
