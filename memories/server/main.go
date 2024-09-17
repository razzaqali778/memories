package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"your_project/routes"
)

var postCollection *mongo.Collection

func main() {
	// MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://js_mastery:123123123@practice.jto9p.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	postCollection = client.Database("test").Collection("posts")

	// Set up Gin server
	router := gin.Default()

	// Middleware for CORS and body parsing
	router.Use(cors.Default())

	// Routes for the posts
	routes.PostRoutes(router, postCollection)

	// Start the server
	port := "5000"
	fmt.Printf("Server Running on Port: http://localhost:%s\n", port)
	log.Fatal(router.Run(":" + port))
}
