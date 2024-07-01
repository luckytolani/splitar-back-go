package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	app := fiber.New()

	Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!555")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	app.Listen(":3000")
}

func Connect() *mongo.Collection {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find .evn
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Get value from .env
	mongoUri := os.Getenv("MONGO_URI")

	// Connect to the database.
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}

	collection := client.Database("testdb").Collection("test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to db")
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	res, err := collection.InsertOne(ctx, bson.D{{Key: "name", Value: "pi"}, {Key: "value", Value: 3.14159}})
	fmt.Println(res.InsertedID)
	// id := res.InsertedID

	return collection
}
