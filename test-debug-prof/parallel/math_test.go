package math

import "testing"

// START1 OMIT
func TestOne(t *testing.T) {
	t.Parallel() // HL
	if Mul(1.0, 2.0) != 2.0 {
		t.Fatal("Double(1, 2) failed") // FIXME: Better log
	}
}

func TestTwo(t *testing.T) {
	t.Parallel() // HL
	if Mul(2.0, 2.0) != 4 {
		t.Fatal("Double(2, 2) failed") // FIXME: Better log
	}
}
// END1 OMIT
