package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient interface {
	// InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*InsertOneResult, error)
}

type MongoDBSaver struct {
	Client mongoClient
}

type History struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	UUID           string             `bson:"uuid"`
	DeviceID       string             `bson:"device_id"`
	PackageCode    string             `bosin:"package_code"`
	Status         string             `bson:"status"` // Status("NEW", "SUCCESS", "FAIL")
	UserID         string             `bson:"uid"`
	SpaceID        string             `bson:"space_id"`
	ProductID      string             `bson:"product_id"`
	SSOID          string             `bson:"ssoid,omitempty"`
	AdditionalInfo string             `bson:"additional_info,omitempty"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      *time.Time         `bson:"updated_at,omitempty"`
	RequestID      string             `bson:"request_id"`
}

func (s *MongoDBSaver) Save(ctx context.Context, history *History) error {
	// item, err := dynamodbattribute.MarshalMap(p)
	// if err != nil {
	// 	return fmt.Errorf("failed to marshal shoutout for storage: %s", err)
	// }

	// input := &dynamodb.PutItemInput{
	// 	Item:      item,
	// 	TableName: aws.String(os.Getenv("TABLE_NAME")),
	// }

	// _, err = s.Client.PutItemWithContext(ctx, input)

	// return err
	// coll := s.Client.Database("ipc_auto_added").Collection("history")
	// _, err := coll.InsertOne(context.TODO(), history)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err := s.Client.InsertOne(ctx, history)
	// return err

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// history := History{
	// 	ID:          primitive.NewObjectID(),
	// 	PackageCode: "GIV50OFF",
	// 	UUID:        "jdwsc2027a1fd34749c0",
	// 	DeviceID:    "eb4337154c015e5c80uiru",
	// 	Status:      "NEW",
	// 	SpaceID:     "202231538",
	// 	UserID:      "az1718348366909EXmeu",
	// 	ProductID:   "tda8yu7ankwszilh",
	// 	CreatedAt:   time.Now(),
	// 	RequestID:   uuid.New().String(),
	// }

	// coll := client.Database("ipc_auto_added").Collection("history")
	// _, err = coll.InsertOne(context.TODO(), history)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mongoSaver := MongoDBSaver{Client: client}
	// if err := mongoSaver.Save(nil, &history); err != nil {
	// 	log.Fatalf("failed to save: %v", err)
	// }

	coll := client.Database("ipc_auto_added").Collection("history")
	filter := bson.D{{"product_id", "tda8yu7ankwszilh"}, {"uuid", "jdwsc2027a1fd34749c0"}}
	var result bson.D
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(result)

	if result == nil {
		fmt.Println("insert to db")
	}

	// coll := client.Database("audo_add_cloud").Collection("movies")
	// title := "Back to the Future"
	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).
	// 	Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
