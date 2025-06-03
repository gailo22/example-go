package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type BucketBasics2 struct {
	S3Client *s3.Client
}

type Presigner2 struct {
	PresignClient *s3.PresignClient
}

// func (basics BucketBasics2) UploadFile(bucketName string, objectKey string, fileName string) error {

// 	// file, err := ioutil.TempFile("dir", "myname.*.png")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer os.Remove(file.Name())

// 	// fmt.Println(file.Name()) // For example "dir/myname.054003078.bat"

// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
// 	} else {
// 		defer file.Close()
// 		_, err = basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
// 			Bucket: aws.String(bucketName),
// 			Key:    aws.String(objectKey),
// 			Body:   file,
// 		})
// 		if err != nil {
// 			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
// 				fileName, bucketName, objectKey, err)
// 		}
// 	}
// 	return err
// }

func (basics BucketBasics2) UploadFile(bucketName string, objectKey string, file *os.File) error {
	defer os.Remove(file.Name())
	_, err := basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
			file.Name(), bucketName, objectKey, err)
	}

	return err
}

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
	fileURL := "https://ty-us-storage30.s3.dualstack.us-west-2.amazonaws.com/da0d77-153971403-jdwsc2027a1fd34749c0/unify/1725933978.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20240910T020736Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUTPMUJJJZZBR5PEK%2F20240910%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Expires=3599&X-Amz-Signature=21439610de6222b6cc496da388e88d4abddfc0109fdddd23efd798096107d8ab"

	// Decrypt file from URL and save to destination
	// tempFile, err := decryptFile(fileKey, fileURL, "2022032511-123.jpg")
	// if err != nil {
	// 	fmt.Println("Error decrypting file:", err)
	// }

	// savePath := "/Users/montree/dev/tmp/go-exercise/example/s3/downloaded_encrypted_image.jpg"
	tmpDecryptedFile := "downloaded_encrypted_image.jpg"
	tmpFilePath, err := downloadEncryptedImage(fileURL, tmpDecryptedFile)
	if err != nil {
		log.Fatalf("Failed to download image: %v", err)
	}

	log.Println("Image downloaded successfully:", tmpFilePath)

	in, err := os.Open(tmpFilePath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer in.Close()
	defer os.Remove(tmpFilePath)

	// decryptedFileName := "/Users/montree/dev/tmp/go-exercise/example/s3/2022032511-3.jpg"
	decryptedFileName := "2022032511-3.jpg"
	f, err := decryptFile3(fileKey, in, decryptedFileName)
	if err != nil {
		fmt.Println("Error decrypting file:", err)
	} else {
		fmt.Println("File decrypted successfully:", f.Name())
	}

	defer os.Remove(f.Name())

	decryptedFilePath := f.Name()

	bucket := "my-bughoong-bucket"
	key := fmt.Sprintf("snapshots/2022032511-%v.jpg", time.Now().Unix()) // uuid + timestamp.jpg

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	bucketOps := BucketBasics2{client}
	presignClient := s3.NewPresignClient(client)

	decryptedFile, err := os.Open(decryptedFilePath)
	err = bucketOps.UploadFile(bucket, key, decryptedFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("update file %v successful", decryptedFilePath)

	presigner := Presigner2{
		PresignClient: presignClient,
	}

	request, err := presigner.PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(5 * time.Minute)
	})

	fmt.Println("presigned url:", request.URL)
}

func downloadEncryptedImage(url, filepath string) (string, error) {
	// Send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create a new file on the local system
	file, err := os.CreateTemp("", filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the response body (image data) to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

// Function to decrypt a file from a URL and save it to a destination
func decryptFile(key, fileURL, destPath string) (*os.File, error) {
	// Get file from URL
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create destination file
	// destFile, err := os.Create(destPath)
	// if err != nil {
	// 	return err
	// }
	// defer destFile.Close()
	destFile, err := os.CreateTemp("", destPath)
	if err != nil {
		log.Fatal(err)
	}
	// defer os.Remove(destFile.Name())

	// Read IV from input stream
	iv := make([]byte, ivLength)
	if _, err := io.ReadFull(resp.Body, iv); err != nil {
		return nil, err
	}

	// Initialize AES cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt data and write to destination file
	buf := make([]byte, bufferSize)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		mode.CryptBlocks(buf[:n], buf[:n])
		if _, err := destFile.Write(buf[:n]); err != nil {
			return nil, err
		}
	}

	return destFile, nil
}

func decryptFile3(key string, in io.Reader, destFile string) (*os.File, error) {
	// Open the destination file for writing
	out, err := os.CreateTemp("", destFile)
	if err != nil {
		return nil, err
	}
	defer out.Close()

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
