package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"your_project/controllers"
)

func PostRoutes(router *gin.Engine, postCollection *mongo.Collection) {
	// Pass the collection to the controller functions
	router.GET("/posts", func(c *gin.Context) {
		controllers.GetPosts(c, postCollection)
	})

	router.POST("/posts", func(c *gin.Context) {
		controllers.CreatePost(c, postCollection)
	})

	router.GET("/posts/:id", func(c *gin.Context) {
		controllers.GetPost(c, postCollection)
	})

	router.PATCH("/posts/:id", func(c *gin.Context) {
		controllers.UpdatePost(c, postCollection)
	})

	router.DELETE("/posts/:id", func(c *gin.Context) {
		controllers.DeletePost(c, postCollection)
	})

	router.PATCH("/posts/:id/like", func(c *gin.Context) {
		controllers.LikePost(c, postCollection)
	})
}
