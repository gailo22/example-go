package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type Presigner struct {
	PresignClient *s3.PresignClient
}

func main() {

	bucketName := "my-bughoong-bucket"
	objectKey := "snapshots/Donuts.png"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	s3client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3client)

	presigner := Presigner{
		PresignClient: presignClient,
	}

	request, err := presigner.PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(5 * time.Minute)
	})

	fmt.Println("presigned url:", request.URL)

}
