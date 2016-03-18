package mymath

import "testing"

func TestSum(t *testing.T) {
	r, err := Sum(1, 3)
	if err != nil {
		t.Fatal(err)
	}

	if r != 4 {
		t.Fatalf("expected %d was %d", 4, r)
	}
}
