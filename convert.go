package convert

import (
	"fmt"
	"io"
	"strings"
)

type Printer struct {
	Print func(w io.Writer, a ...any) (n int, err error)
}

type option func(*Printer)

func WithPrinter(print func(w io.Writer, a ...any) (n int, err error)) option {
	return func(p *Printer) {
		p.Print = print
	}
}

func NewPrinter(opts ...option) *Printer {
	p := &Printer{Print: fmt.Fprint}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *Printer) String(w io.Writer) string {
	b := new(strings.Builder)
	p.Print(b, w)
	return b.String()
}

func (p *Printer) Stringln(w io.Writer) string {
	b := new(strings.Builder)
	p.Print(b, w)
	return b.String()
}

// String will write the content of an io.Writer to a string builder without a trailing newline.
func String(w io.Writer) string {
	p := NewPrinter()
	return p.String(w)
}

// Stringln will write the content of an io.Writer to a string builder with a trailing newline.
func Stringln(w io.Writer) string {
	p := NewPrinter(WithPrinter(fmt.Fprintln))
	return p.String(w)
}
