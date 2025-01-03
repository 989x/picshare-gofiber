package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GeneratePublicID generates a random public_id in the format xxyyxxyy
func GeneratePublicID() string {
	b := make([]byte, 4) // 4 bytes = 8 characters in hex
	rand.Read(b)
	hexID := hex.EncodeToString(b)
	return fmt.Sprintf("%c%c%c%c%c%c%c%c",
		hexID[0], hexID[1], hexID[2],
		'a'+(hexID[3]-'0')%6, // Convert hex to 'a-f'
		hexID[4], hexID[5], hexID[6],
		'a'+(hexID[7]-'0')%6, // Convert hex to 'a-f'
	)
}
