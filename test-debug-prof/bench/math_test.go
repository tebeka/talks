package math

import "testing"

// START1 OMIT
func BenchmarkDiv(b *testing.B) { // HL
	for i := 0; i < b.N; i++ { // go test will increase N to get good timing 
		Div(10.0, 3.0)
	}
}
// END1 OMIT

