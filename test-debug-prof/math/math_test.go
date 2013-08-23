package math

// START1 OMIT
import "testing" // HL

func TestMul(t *testing.T) {
	a, b, expected := 2.0, 3.0, 6.0
	result := Mul(a, b)
	if result != expected {
		t.Fatalf("%v * %v -> %v (expected %v)", a, b, result, expected) // HL
	}
}
// END1 OMIT

// START2 OMIT
func TestDiv(t *testing.T) {
	a, b, expected := 6.0, 3.0, 3.0
	result := Div(a, b)
	if result != expected {
		t.Fatalf("%v / %v -> %v (expected %v)", a, b, result, expected) // HL
	}
}
// END2 OMIT
