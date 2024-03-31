package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
    Name         string `bson:"name"`
    Qty          int    `bson:"qty"`
    Product      string `bson:"product"`
    Availability bool   `bson:"availability"`
}

var collection *mongo.Collection

func setupDatabase() {
    // MongoDB connection URI
    connectionURI := "mongodb://localhost:27017"
    clientOptions := options.Client().ApplyURI(connectionURI)

    // Connect to MongoDB
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Set the global collection variable
    database := client.Database("local")
    collection = database.Collection("sample_db")
}

func CreateItem(c *fiber.Ctx) error {
    var item Item
    if err := c.BodyParser(&item); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    insertResult, err := collection.InsertOne(context.Background(), item)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "Item created", "inserted_id": insertResult.InsertedID})
}

func GetItem(c *fiber.Ctx) error {
    name := c.Params("name")
    filter := bson.D{{Key: "name", Value: name}}

    var result Item
    err := collection.FindOne(context.Background(), filter).Decode(&result)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
    }

    return c.JSON(result)
}

func GetItems(c *fiber.Ctx) error {
    // Query parameters
    limit, _ := strconv.Atoi(c.Query("limit", "10"))
    page, _ := strconv.Atoi(c.Query("page", "1"))
    sort := c.Query("sort", "name")

    // Calculate skip value for pagination
    skip := (page - 1) * limit
    if skip < 0 {
        skip = 0
    }

    // Options for find operation (including pagination and sorting)
    findOptions := options.Find()
    findOptions.SetLimit(int64(limit))
    findOptions.SetSkip(int64(skip))
    findOptions.SetSort(bson.D{{Key: sort, Value: 1}}) // 1 for ascending order, -1 for descending order

    // Find documents
    cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    defer cursor.Close(context.Background())

    // Decode documents into a slice
    var results []Item
    if err := cursor.All(context.Background(), &results); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(results)
}

func UpdateItem(c *fiber.Ctx) error {
    name := c.Params("name")
    filter := bson.D{{Key: "name", Value: name}}

    var updateData bson.M // Use a map for unmarshaling JSON
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    update := bson.D{{Key: "$set", Value: updateData}}

    updateResult, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    if updateResult.ModifiedCount == 0 {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
    }

    return c.JSON(fiber.Map{"message": "Item updated"})
}

func DeleteItem(c *fiber.Ctx) error {
    name := c.Params("name")
    filter := bson.D{{Key: "name", Value: name}}

    deleteResult, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    if deleteResult.DeletedCount == 0 {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
    }

    return c.JSON(fiber.Map{"message": "Item deleted"})
}git branch -M main

func main() {
    app := fiber.New()

    // Setup MongoDB connection
    setupDatabase()

    // Routes
    app.Post("/inventory", CreateItem)
    app.Get("/inventory/:name", GetItem)
    app.Put("/inventory/:name", UpdateItem)
    app.Delete("/inventory/:name", DeleteItem)
    app.Get("/inventories", GetItems)

    // Start the server
    port := 3000
    fmt.Printf("Server listening on port %d...\n", port)
    log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}