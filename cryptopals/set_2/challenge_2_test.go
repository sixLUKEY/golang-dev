package xorcombo

import (
	"encoding/hex"
	"testing"
)

func TestXORBuffers(t *testing.T) {
	// Example buffers
	buf1, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		t.Errorf("Unable to decode string")
	}

	buf2, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		t.Errorf("Unable to decode string")
	}

	expectedResult := "746865206b696420646f6e277420706c6179"

	// XOR the buffers
	actualResult, err := XORBuffers(buf1, buf2)
	if err != nil {
		t.Errorf("Error:")
	}

	if hex.EncodeToString(actualResult) != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, actualResult)
	}

	invalidXORstring := "invalidXOR"
	alsoInvalidXor := "alsoInvalid"

	_, err = XORBuffers([]byte(invalidXORstring), []byte(alsoInvalidXor))
	if err == nil {
		t.Errorf("Expected an error for invalid XOR string, but got nil")
	}
}
