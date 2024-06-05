package HexToBase64

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexString string) (string, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	base64String := base64.RawStdEncoding.EncodeToString(bytes)
	return base64String, nil
}
