package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	want := 55
	got := FibonacciFor(10)

	if got != want {
		t.Errorf("FibonacciFor(10) = %d; want %d", got, want)
	}

}
