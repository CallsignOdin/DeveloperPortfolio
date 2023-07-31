package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Serve static files
	router.Static("/css", "static/css")
	router.Static("/scripts", "static/scripts")

	// Load the HTML template
	router.LoadHTMLGlob("templates/*")

	// Define a route for the root path
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)

	})

	// Start the server on port 8080
	router.Run(":8080")
}
