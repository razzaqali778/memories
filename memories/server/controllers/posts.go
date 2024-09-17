package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"your_project/models"  // Import the models package
)

// GetPosts retrieves all posts
func GetPosts(c *gin.Context, postCollection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := postCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var posts []models.Post  // Use the Post model
	if err = cursor.All(ctx, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPost retrieves a single post by its ID
func GetPost(c *gin.Context, postCollection *mongo.Collection) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var post models.Post  // Use the Post model
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := postCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// CreatePost creates a new post
func CreatePost(c *gin.Context, postCollection *mongo.Collection) {
	var post models.Post  // Use the Post model
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.ID = primitive.NewObjectID()
	post.CreatedAt = time.Now()
	post.LikeCount = 0

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := postCollection.InsertOne(ctx, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// UpdatePost updates an existing post
func UpdatePost(c *gin.Context, postCollection *mongo.Collection) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var updatedPost models.Post  // Use the Post model
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": updatedPost,
	})
	if err != nil || result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// DeletePost deletes a post by its ID
func DeletePost(c *gin.Context, postCollection *mongo.Collection) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := postCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// LikePost increases the like count of a post
func LikePost(c *gin.Context, postCollection *mongo.Collection) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var post models.Post  // Use the Post model
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := postCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	_, err = postCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$inc": bson.M{"likeCount": 1},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}
