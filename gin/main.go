package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func main() {
	router := gin.Default()

	// Global middleware
	router.Use(globalMiddleware)

	// Grouping routes
	api := router.Group("/api")
	{
		api.GET("/users/:id", getUser)
		api.POST("/users", createUser)
	}
	router.GET("/", hello)
	router.Run(":8080")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}

// Global middleware example
func globalMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
}

// Route handler with parameterized URL
func getUser(c *gin.Context) {
	id := c.Param("id")
	// Perform user retrieval based on the provided ID (not implemented in this example)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not provided"})
		return
	}
	// Simulate response for demonstration purposes
	c.JSON(http.StatusOK, gin.H{"user_id": id})
}

// Route handler for handling POST request with JSON data
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		// Validation error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save the user to the database (not implemented in this example)
	// Simulate successful response for demonstration purposes
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
