package encrypting

const sampleSeed = 42
const sampleShift uint = 42

var sampleData = []byte{3, 14, 15, 92, 6, 5, 4}

func caesarEncode(data []byte, shift uint) []byte {
	result := make([]byte, len(data))
	for index, value := range sampleData {
		result[index] = byte((uint(value) + shift) % (1 << 8))
	}

	return result
}
