package document

import (
	"fmt"
)

type Element interface {
	fmt.Stringer
	ToHTML() string
}
