package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ")
	}

	MONGODB_URL := os.Getenv("MONGODB_URL")

	clientOption := options.Client().ApplyURI(MONGODB_URL)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	fmt.Println("Connected to MONGODB ATLAS")

	collection = client.Database("react-go-tutorial").Collection("todos")

	app := fiber.New()
	app.Get("/api/todos", GetTodos)
	app.Post("/api/todos", AddTodos)
	app.Patch("/api/todos/:id", UpdateTodos)
	app.Delete("/api/todos/:id", DeleteTodos)

	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))
}

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.Status(http.StatusOK).JSON(todos)
}

func AddTodos(c *fiber.Ctx) error {

	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": "Todo body is required"})
	}
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(http.StatusCreated).JSON(todo)

}

func UpdateTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": "Invalid object ID"})
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"success": true})
}

func DeleteTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": "Invalid object ID"})
	}
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"success": true})
}
