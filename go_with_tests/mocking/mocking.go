package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%v\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(out, "Go!")
}
