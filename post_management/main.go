package main

import (
	"fmt"
	"post_management/database"
	"post_management/handlers"
	"post_management/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello, World!")

	// Initialize the Gin router
	router := gin.Default()

	// Initialize the database connection
	database.InitDB()

	// Define the routes and their handlers
	router.POST("/signup", handlers.CreateUser)
	router.POST("/login", handlers.Login)
	router.POST("/post", middlewares.CheckAuth(), handlers.CreatePost)
	router.GET("/posts", middlewares.CheckAuth(), handlers.GetPosts)
	router.DELETE("/post/:id", middlewares.CheckAuth(), handlers.DeletePost)
	router.POST("/comment/:post_id", middlewares.CheckAuth(), handlers.CreateComment)
	router.GET("/comments/:post_id", middlewares.CheckAuth(), handlers.GetComments)
	router.DELETE("/comment/:comment_id", middlewares.CheckAuth(), handlers.DeleteComment)
	router.Run(":8080") // Run on port 8080
	fmt.Println("Server is running on port 8080")
}
