package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// Example input JSON values
	encryptedData := "246c2eae9fd248bb6f42f8124bd7ed78b505fb6e470da9a8e5c43f3bab5340579f53f4ea71e074a77c72f4627ac62fdddf7090e184eec4f64e8af9c476936092a4f1739c05ba7f9d70af088249528a824ed32615d33f6f015b54fc0e3f6cf8501c37ec9e0d779afe0e97b111289ed56e"
	iv := "59a14c07305b34ceff3e4ef7e505a13e"
	// accessKey := "42ee90ce9aed4bbf98c27d75e44bf6c8"
	accessKey := "2b3ca2a1e8314a3895775913444c8091"
	key := accessKey[8:24]

	// Convert hex-encoded values to byte slices
	dataBytes, err := hex.DecodeString(encryptedData)
	if err != nil {
		fmt.Println("Failed to decode encrypted data:", err)
		return
	}

	ivBytes, err := hex.DecodeString(iv)
	if err != nil {
		fmt.Println("Failed to decode IV:", err)
		return
	}

	keyBytes := []byte(key) // Assuming key is a plain string; you may need to adjust it

	// Ensure key length is appropriate for AES (16, 24, or 32 bytes)
	if len(keyBytes) != 16 && len(keyBytes) != 24 && len(keyBytes) != 32 {
		fmt.Println("Invalid key length:", len(keyBytes))
		return
	}

	// Decrypt the data using AES
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Println("Failed to create AES cipher:", err)
		return
	}

	mode := cipher.NewCBCDecrypter(block, ivBytes)
	decrypted := make([]byte, len(dataBytes))
	mode.CryptBlocks(decrypted, dataBytes)

	// Remove padding if necessary (PKCS#7 padding is often used)
	decrypted = pkcs7Unpad(decrypted)

	fmt.Println("Decrypted data:", string(decrypted))
}

// pkcs7Unpad removes padding from decrypted data
func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
