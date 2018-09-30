// +build js

package ulog

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
