package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	message, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		return
	}

	buf1, err := hex.DecodeString("abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()_+:',./<>?'")

	result := make([]byte, len(message))

	for i := 0; i < len(buf1); i++ {
		result[i] = buf1[i] ^ message[i]
	}

	// fmt.Println(message)
	fmt.Println(result)
	fmt.Println(hex.EncodeToString(result))
}
