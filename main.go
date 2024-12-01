package main

import (
	"context"
	"log"
	"os"

	router "myary/routes"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

var (
	mongoClient     *mongo.Client
	firestoreClient *firestore.Client
	uri             string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	uri = os.Getenv("DATABASE_URL")
	if uri == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	if err := connect_to_firestore(); err != nil {
		log.Fatal("Could not connect to Firestore:", err)
	}
}

func main() {
	r := gin.Default()
	db := mongoClient.Database("myary")

	router.SetupRoutes(r, db)

	if err := r.Run(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func connect_to_mongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client

	log.Println("Connected to MongoDB successfully")
	return err
}

func connect_to_firestore() error {
	credentialsFile := "./secret/myary-8d088-firebase-adminsdk-tn42v-83c3f2d74b.json"
	opt := option.WithCredentialsFile(credentialsFile)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	firestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		return err
	}

	log.Println("Connected to Firestore successfully")
	return nil
}
