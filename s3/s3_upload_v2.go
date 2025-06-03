package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type BucketBasics struct {
	S3Client *s3.Client
}

func (basics BucketBasics) UploadFile(bucketName string, objectKey string, fileName string) error {

	// file, err := ioutil.TempFile("dir", "myname.*.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer os.Remove(file.Name())

	// fmt.Println(file.Name()) // For example "dir/myname.054003078.bat"

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		defer file.Close()
		_, err = basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, objectKey, err)
		}
	}
	return err
}

// Here, you can choose the region of your bucket
func main() {
	// region := "ap-southeast-1"
	bucket := "my-bughoong-bucket"
	key := "snapshots/Donuts.png"
	fileName := "Donuts.png"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	bucketOps := BucketBasics{client}

	err = bucketOps.UploadFile(bucket, key, fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("update file %v successful", fileName)

}
