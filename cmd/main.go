package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/GoToDo/database"
	"github.com/ienjir/GoToDo/routes"
)

func main() {
	r := gin.Default()

	// Initialize the database
	database.ConnectDatabase()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	err := r.Run(":8080")
	if err != nil {
		return
	} // Default port is 8080
	fmt.Println("Server started")
}
