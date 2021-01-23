package outliers_test

import (
	"fmt"
	"math/rand"

	"github.com/353solutions/outliers"
)

func Example() {
	o, err := outliers.New("outliers", "detect")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	defer o.Close()

	data := genData()
	indices, err := o.Detect(data)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Println(indices)
	// Output:
	// [7 113 835]
}

func genData() []float64 {
	const size = 1000
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Float64()
	}

	indices := []int{7, 113, 835}
	for _, i := range indices {
		data[i] += 97
	}

	return data
}
