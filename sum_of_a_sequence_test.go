package sequencesum

import "testing"

func TestSequenceSum(t *testing.T) {
	sum := SequenceSum(1, 5, 3)
	want := 5

	if sum != want {
		t.Errorf("got %d want %d", sum, want)
	}
}
