package dbmanager

import (
	"context"
	"fmt"
	"os"

	// "github.com/joho/godotenv"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var connectionString string
const dbName = "url_shortner"
const collectionName = "url_shortner_collection"

var collection *mongo.Collection

type Document struct{
	ShortLink string `bson:"key"`
	RedirectURL string `bson:"value"`
}

func Init() {
	godotenv.Load()
	connectionString = os.Getenv("MONGO_STRING")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err!= nil {
        panic(err)
    }
	fmt.Println("Connected to MongoDB")
	collection = client.Database(dbName).Collection(collectionName)
}


func InsertOne(actualURL string,shortURL string)  {
	doc := Document{
        ShortLink: shortURL,
        RedirectURL: actualURL,
    }
	id,err:=collection.InsertOne(context.TODO(),doc)
	if err!= nil {
        panic(err)
    }
	fmt.Println("inserted with ID",id)
}

func FindOne(shortURL string) string {
	filter := bson.M{"key":shortURL}
	var data Document
	err := collection.FindOne(context.TODO(),filter).Decode(&data)
	if err!= nil {
        fmt.Println("No document found with given ID")
        return ""
    }
	return data.RedirectURL
}
