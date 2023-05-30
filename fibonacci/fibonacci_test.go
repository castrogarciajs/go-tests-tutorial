package fibonacci

import (
	"testing"
)

func TestFibonacciFor(t *testing.T) {
	want := 55
	got := FibonacciFor(10)

	if got != want {
		t.Errorf("FibonacciFor(10) = %d; want %d", got, want)
	}

}

func TestFibonacciRec(t *testing.T) {

	want := 55
	got := FinonacciRec(10)

	if got != want {
		t.Errorf("Se esperaba  %d, se obtuvo %d", want, got)
	}

}

func TestFibonacciRecMemo(t *testing.T) {
	want := 55
	got := fibonacciRecMemo(10)

	if got != want {
		t.Errorf("Se esperaba %d, se obtuvo %d", want, got)
	}
}
