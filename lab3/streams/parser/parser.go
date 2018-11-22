package parser

import (
	"github.com/pkg/errors"
	"ood/lab3/streams/decorator"
	"strconv"
)

type ParsedArgs struct {
	RDecorator     decorator.ReaderDecorator
	WDecorator     decorator.WriterDecorator
	InputFileName  string
	OutputFileName string
}

type Parser interface {
	Parse([]string) (*ParsedArgs, error)
}

type parser struct{}

func New() Parser {
	return new(parser)
}

func (p parser) Parse(args []string) (*ParsedArgs, error) {
	if len(args) < 2 {
		return nil, errors.New("Not enough args")
	}

	rArgs, wArgs := p.split(args[:len(args)-2])
	rDecorator, err := p.parseReaderArgs(rArgs)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse reader args")
	}

	wDecorator, err := p.parseWriterArgs(wArgs)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse writer args")
	}

	return &ParsedArgs{
		RDecorator:     rDecorator,
		WDecorator:     wDecorator,
		InputFileName:  args[len(args)-2],
		OutputFileName: args[len(args)-1],
	}, nil
}

func (p parser) parseReaderArgs(args []string) (decorator.ReaderDecorator, error) {
	var decorators []decorator.ReaderDecorator
	for index := 0; index < len(args); index++ {
		println(args[index])
		switch args[index] {
		case "--decompress":
			decorators = append(decorators, new(decorator.DecompressionReaderDecorator))
		case "--decrypt":
			key, err := p.parseKey(args, index+1)
			if err != nil {
				return nil, errors.Wrap(err, "Failed to parse key")
			}
			decorators = append(decorators, decorator.NewDecryptionReaderDecorator(key))
			index++
		default:
			return nil, errors.Errorf("Unexpected arg: %s", args[index])
		}
	}
	return decorator.CompositeReaderDecorator(decorators), nil
}

func (p parser) parseWriterArgs(args []string) (decorator.WriterDecorator, error) {
	var decorators []decorator.WriterDecorator
	for index := 0; index < len(args); index++ {
		println(args[index])
		switch args[index] {
		case "--compress":
			decorators = append(decorators, new(decorator.CompressionWriterDecorator))
		case "--encrypt":
			key, err := p.parseKey(args, index+1)
			if err != nil {
				return nil, errors.Wrap(err, "Failed to parse key")
			}
			decorators = append(decorators, decorator.NewEncryptionWriterDecorator(key))
			index++
		default:
			return nil, errors.Errorf("Unexpected arg: %s", args[index])
		}
	}
	return decorator.CompositeWriterDecorator(decorators), nil
}

func (parser) parseKey(args []string, index int) (int64, error) {
	if index >= len(args) {
		return 0, errors.Errorf("Expected key value after %s arg", args[index-1])
	}

	key, err := strconv.ParseInt(args[index], 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "Failed to parse int `%d`", args[index])
	}

	return key, nil
}

func (parser) split(args []string) ([]string, []string) {
	for index, val := range args {
		if val == "--compress" || val == "--encrypt" {
			return args[:index], args[index:]
		}
	}
	return args, []string{}
}
