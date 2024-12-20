package mocking

import (
	"bytes"
	"testing"
	"reflect"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Test countdown", func (t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := SpySleeper{}
		Countdown(&buffer, &spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("expected %q sleep calls, got %q", 3, spySleeper.Calls)
		}
	})
	t.Run("Test order of countdown", func (t *testing.T) {
		spySleepPrinter := SpyCountdownOperations{}
		Countdown(&spySleepPrinter, &spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,	
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}