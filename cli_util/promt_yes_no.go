package cli_util

import (
	"bufio"
	"fmt"
	"os"
)

func PromtYesNo(question string) bool {
	fmt.Print(question + " (y/n) ")
	stdinReader := bufio.NewReader(os.Stdin)
	var ch rune
	var err error
	for err == nil && (ch == rune(0) || ch == '\n') {
		ch, _, err = stdinReader.ReadRune()
	}
	return ch == 'y' || ch == 'Y'
}
