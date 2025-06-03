package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Here, you can choose the region of your bucket
func main() {
	region := "ap-southeast-1"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}
	svc := s3.New(sess)

	bucket := "my-bughoong-bucket"
	filePath := "Donuts.png"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		return
	}
	defer file.Close()

	key := "Donuts.png"

	// Read the contents of the file into a buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		return
	}

	// This uploads the contents of the buffer to S3
	// _, err = svc.PutObject(&s3.PutObjectInput{
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String(key),
	// 	Body:   bytes.NewReader(buf.Bytes()),
	// })

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
	})

	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	str, err := req.Presign(15 * time.Minute)

	log.Println("The URL is:", str, " err:", err)

	fmt.Println("File uploaded successfully!!!")
}
