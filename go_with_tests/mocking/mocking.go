package mocking

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3\n\t2\n\t1\n\tGo!")
}
