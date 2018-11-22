package encryption

const (
	sampleLen        = 1000
	sampleSeed       = 42
	sampleShift uint = 42
)

func getSampleData() []byte {
	data := []byte{3, 14, 15, 92, 6, 5, 4}
	result := make([]byte, len(data))
	copy(result, data)
	return result
}

func caesarEncode(data []byte, shift uint) []byte {
	result := make([]byte, len(data))
	for index, value := range getSampleData() {
		result[index] = byte((uint(value) + shift) % (1 << 8))
	}

	return result
}
