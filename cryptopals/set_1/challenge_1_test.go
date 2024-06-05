package HexToBase64

import (
	"testing"
)

func TextHexToBase(t *testing.T) {
	hextString := "48656c6c6f20576f726c64"
	expectedBase64 := "SGVsbG8gV29ybGQ="
	actualBase64, err := HexToBase64(hextString)
	if err != nil {
		t.Errorf("Error converting hex to base64: %v", err)
	}
	if actualBase64 != expectedBase64 {
		t.Errorf("Expected %s, but got %s", expectedBase64, actualBase64)
	}

	invalidHexString := "invalidHex"
	_, err = HexToBase64(invalidHexString)
	if err == nil {
		t.Errorf("Expected an error for invalid hex string, but got nil")
	}

}
