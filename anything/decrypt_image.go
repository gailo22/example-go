package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Constants for encryption
const (
	defaultAlgorithm     = "AES"
	defaultFullAlgorithm = "AES/CBC/PKCS5Padding"
	keyLength            = 16 // AES-128 requires a 16-byte key
	ivLength             = 16 // Initialization vector for AES
	bufferSize           = 1024
)

func main() {
	fileKey := "01063c5c3d684933"
	// fileURL := "https://ty-us-storage30.s3.dualstack.us-west-2.amazonaws.com/da0d77-153971403-jdwsc2027a1fd34749c0/unify/1725900876.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20240909T165641Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUTPMUJJJZZBR5PEK%2F20240909%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Expires=3599&X-Amz-Signature=8a4ff97a6620f0ebcba646e48a9867a3132f183c3e182d1f096a503f0d5eccf0"

	// Decrypt file from URL and save to destination
	// err := decryptFile(fileKey, fileURL, "/Users/montree/tmp/2022032511-2.jpg")
	// if err != nil {
	// 	fmt.Println("Error decrypting file:", err)
	// }

	in, err := os.Open("/Users/montree/tmp/image_encrypt.jpeg")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer in.Close()

	_, err = decryptFile3(fileKey, in, "/Users/montree/tmp/2022032511-3.jpg")
	if err != nil {
		fmt.Println("Error decrypting file:", err)
	} else {
		fmt.Println("File decrypted successfully")
	}
}

// Function to decrypt a file from a URL and save it to a destination
func decryptFile(key, fileURL, destPath string) error {
	// Get file from URL
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	in := resp.Body

	// Create destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Read IV from input stream
	iv := make([]byte, ivLength)
	if _, err := io.ReadFull(in, iv); err != nil {
		return err
	}
	if len(iv) != 16 {
		return errors.New("invalid IV length")
	}

	// Skip header bytes
	_, err = io.ReadFull(in, make([]byte, 44))
	if err != nil {
		return err
	}

	// Initialize AES cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt data and write to destination file
	buf := make([]byte, bufferSize)
	for {
		// n, err := in.Read(buf)
		n, err := io.ReadFull(in, buf)
		// n, err := io.ReadAll(io.LimitReader(in, bufferSize))
		// n, err := io.ReadAtLeast(in, buf, len(buf))
		fmt.Println("n:", n)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		src := buf[:n]
		fmt.Println("src len:", len(src))
		// mode.CryptBlocks(buf[:n], src)
		// blockMode := cipher.NewCBCDecrypter(block, iv)
		origData := make([]byte, len(src))
		mode.CryptBlocks(origData, src)
		if _, err := destFile.Write(origData); err != nil {
			return err
		}
	}

	return nil
}

func decryptFile2(key string, fileURL string, destPath string) error {
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	in := resp.Body

	// Create destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Read and verify the IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.CopyN(io.Discard, in, 4); err != nil {
		return err
	}
	if _, err := io.ReadFull(in, iv); err != nil {
		return fmt.Errorf("iv length error: %v", err)
	}
	if _, err := io.CopyN(io.Discard, in, 44); err != nil {
		return err
	}

	// Convert key to bytes
	// keyBytes, err := hex.DecodeString(key)
	// if err != nil {
	// 	return fmt.Errorf("invalid key: %v", err)
	// }

	// Initialize AES cipher
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return err
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt file content
	buf := make([]byte, 1024)
	var n = 0
	for {
		// n, err := in.Read(buf)
		n, err = io.ReadFull(in, buf)
		fmt.Println("n:", n)
		// if n < 1024 {
		// 	if _, err := io.CopyN(io.Discard, in, int64(n)); err != nil {
		// 		return err
		// 	}
		// }
		if err != nil && err != io.EOF {
			fmt.Println("eeeee:", err)
			n, err = in.Read(buf)
			// return err
		}
		if n == 0 {
			break
		}

		// if n < 1024 {
		// 	break
		// }

		// Write the decrypted data
		// decrypted := make([]byte, n)
		// mode.CryptBlocks(decrypted, buf[:n])
		// destFile.Write(decrypted)
		src := buf[:n]
		fmt.Println("src len:", len(src))
		// fmt.Println("blockSize:", mode.BlockSize())

		// if len(src)%mode.BlockSize() != 0 {
		// 	continue
		// }
		// mode.CryptBlocks(buf[:n], src)
		// blockMode := cipher.NewCBCDecrypter(block, iv)
		origData := make([]byte, len(src))
		mode.CryptBlocks(origData, src)
		origData = PKCS5UnPadding2(origData)
		if _, err := destFile.Write(origData); err != nil {
			return err
		}
	}

	return nil
}

func decryptFile3(key string, in io.Reader, destFile string) (*os.File, error) {
	// Open the destination file for writing
	out, err := os.Create(destFile)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	// Ensure the parent directory exists
	// if err := os.MkdirAll(destFile, 0755); err != nil {
	// 	return nil, err
	// }

	// Read and verify the IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.CopyN(io.Discard, in, 4); err != nil {
		return nil, err
	}
	if _, err := io.ReadFull(in, iv); err != nil {
		return nil, fmt.Errorf("iv length error: %v", err)
	}
	if _, err := io.CopyN(io.Discard, in, 44); err != nil {
		return nil, err
	}

	// Convert key to bytes
	// keyBytes, err := hex.DecodeString(key)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid key: %v", err)
	// }

	// Initialize AES cipher
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt file content
	buf := make([]byte, 1024)
	for {
		n, err := in.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}

		// Write the decrypted data
		decrypted := make([]byte, n)
		mode.CryptBlocks(decrypted, buf[:n])
		out.Write(decrypted)
	}

	return out, nil
}

func PKCS5Padding2(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding2(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// Function to encrypt a file
func encryptFile(key string, originFilePath string, destPath string) error {
	// Open origin file
	originFile, err := os.Open(originFilePath)
	if err != nil {
		return err
	}
	defer originFile.Close()

	// Create destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Generate random IV
	iv := make([]byte, ivLength)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	// Initialize AES cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	mode := cipher.NewCBCEncrypter(block, iv)

	// Write IV to destination file
	if _, err := destFile.Write(iv); err != nil {
		return err
	}

	// Encrypt data and write to destination file
	buf := make([]byte, bufferSize)
	for {
		n, err := originFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		mode.CryptBlocks(buf[:n], buf[:n])
		if _, err := destFile.Write(buf[:n]); err != nil {
			return err
		}
	}

	return nil
}
