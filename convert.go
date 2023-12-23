package convert

import (
	"fmt"
	"io"
	"strings"
)

// String will write the content of an io.Writer to a string builder without a trailing newline.
func String(w io.Writer) string {
	b := new(strings.Builder)
	fmt.Fprint(b, w)
	return b.String()
}

// Stringln will write the content of an io.Writer to a string builder with a trailing newline.
func Stringln(w io.Writer) string {
	b := new(strings.Builder)
	fmt.Fprintln(b, w)
	return b.String()
}
