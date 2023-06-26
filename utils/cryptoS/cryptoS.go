package cryptos

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Generate Crypto
func Generate(input string, secretKey string) string {
	fmt.Println("Input", input)
	// secretKey := setting.AppSetting.CryptoKey

	// Create a new HMAC hasher with SHA256
	hasher := hmac.New(sha256.New, []byte(secretKey))

	// Write the message to the hasher
	hasher.Write([]byte(input))

	// Get the resulting hash value
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecial string
	hashString := hex.EncodeToString(hash)
	return hashString
}
