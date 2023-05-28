package firsttest

import "testing"

func TestSum(t *testing.T) {
	want := 5
	got := Sum(3, 2)

	if got != want {
		t.Logf("Error: Se esperaba %d, se obtuvo: %d", want, got)
		t.Fail()
	}

}

func TestCum(t *testing.T) {
	want := 6
	got := Cum(1, 2, 3)

	if got != want {
		t.Logf("Error: Se esperaba %d, se obtuvo: %d", want, got)
		t.Fail()
	}
}
