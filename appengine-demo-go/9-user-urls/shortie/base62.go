package shortie

const (
	base62Alphabet = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// base62Encode encodes a number to a base62 string representation.
func base62Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	arr := []uint8{}
	base := uint64(len(base62Alphabet))

	for num > 0 {
		rem := num % base
		num = num / base
		arr = append(arr, base62Alphabet[rem])
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}
