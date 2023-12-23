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
		name     string
		function func(io.Writer) string
		printer  convert.Printer
		want     string
	}{
		{
			name:     "Content will not include a trailing newline.",
			function: convert.String,
			printer:  convert.Printer{Print: fmt.Fprint},
			want:     "Hello, world",
		},
		{
			name:     "Content will include a trailing newline",
			function: convert.Stringln,
			printer:  convert.Printer{Print: fmt.Fprintln},
			want:     "Hello, world\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			writer := new(bytes.Buffer)
			printer := convert.WithPrinter(tt.printer.Print)
			p := convert.NewPrinter(printer)
			p.Print(writer, "Hello, world")
			got := tt.function(writer)

			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
