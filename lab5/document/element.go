package document

import (
	"fmt"
)

type element interface {
	fmt.Stringer
	htmlFormatAcceptor
}
