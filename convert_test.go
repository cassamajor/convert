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
		function  func(io.Writer) string
		printType func(w io.Writer, a ...any) (n int, err error)
		want      string
	}{
		{
			name:      "Content will not include a trailing newline.",
			function:  convert.String,
			printType: fmt.Fprint,
			want:      "Hello, world",
		},
		{
			name:      "Content will include a trailing newline",
			function:  convert.Stringln,
			printType: fmt.Fprintln,
			want:      "Hello, world\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			writer := new(bytes.Buffer)
			tt.printType(writer, "Hello, world")
			got := tt.function(writer)

			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
