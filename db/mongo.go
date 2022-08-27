package db

import (
	"context"
	"encoding/json"
	//	"encoding/json"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var client *mongo.Client

func Setup() {
	log.Println("db setup called")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	url := os.Getenv("MONGODB_URI")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(url).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("db setup completed")
}

func AddTask(name string) {
	coll := client.Database("mongodo").Collection("tasks")
	doc := bson.D{{"name", name}}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
}

func ListTasks() {

}

func DeleteTask() {

}

func FindTask(name string) {
	coll := client.Database("mongodo").Collection("tasks")

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Printf("No task found with name %s", name)
	}
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", jsonData)
}
