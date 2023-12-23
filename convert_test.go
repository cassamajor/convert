package convert_test

import (
	"bytes"
	"fmt"
	"github.com/cassamajor/convert"
	"io"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		printFunc func(w io.Writer, a ...any) (n int, err error)
		want      string
	}{
		{
			name:      "Content will not include a trailing newline.",
			printFunc: fmt.Fprint,
			want:      "Hello, world",
		},
		{
			name:      "Content will include a trailing newline",
			printFunc: fmt.Fprintln,
			want:      "Hello, world\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			writer := new(bytes.Buffer)
			printer := convert.WithPrinter(tt.printFunc)
			p := convert.NewPrinter(printer)
			p.Print(writer, "Hello, world") // Populate the writer using the specified print function

			if got := p.String(writer); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
