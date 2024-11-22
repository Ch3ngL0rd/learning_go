package mocking

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%v\n", i)
	}
	fmt.Fprintf(out, "Go!")
}
