package xorcombo

import (
	"errors"
)

// XORBuffers takes two byte slices of equal length and returns their XOR combination
func XORBuffers(buf1, buf2 []byte) ([]byte, error) {
	// check if the lengths of the buffers are equal
	if len(buf1) != len(buf2) {
		return nil, errors.New("Buffers have the same length")
	}

	// Create a slice to hold the result
	result := make([]byte, len(buf1))

	// XOR each byte from the input buffers
	for i := 0; i < len(buf1); i++ {
		result[i] = buf1[i] ^ buf2[i]
	}

	return result, nil
}
