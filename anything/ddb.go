package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ddbClient interface {
	PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error)
}

type DynamoDBSaver struct {
	Client ddbClient
}

type Person struct {
	Name string
}

func (s *DynamoDBSaver) Save(ctx context.Context, p *Person) error {
	item, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return fmt.Errorf("failed to marshal shoutout for storage: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = s.Client.PutItemWithContext(ctx, input)

	return err
}

func main() {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		log.Fatalf("failed to create AWS session: %v", err)
	}

	client := dynamodb.New(sess)

	ddbSaver := DynamoDBSaver{Client: client}
	if err := ddbSaver.Save(nil, &Person{Name: "Johnny"}); err != nil {
		log.Fatalf("failed to save: %v", err)
	}

}
