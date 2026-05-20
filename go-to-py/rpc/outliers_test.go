package outliers

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func genData() ([]float64, []int) {
	const size = 1000
	data := make([]float64, size)
	for i := range size {
		data[i] = rand.Float64()
	}

	indices := []int{7, 113, 835}
	for _, i := range indices {
		data[i] += 97
	}

	return data, indices
}

func TestDetect(t *testing.T) {
	require := require.New(t)

	o, err := New()
	require.NoError(err, "new")
	defer o.Close()

	data, indices := genData()

	out, err := o.Detect(data)
	require.NoError(err, "detect")
	require.Equal(indices, out, "outliers")
}
