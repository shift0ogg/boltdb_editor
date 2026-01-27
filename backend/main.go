package main

import (
	"boltdb-editor/backend/handlers"
	"boltdb-editor/backend/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database service
	dbService := services.NewDBService()

	// Initialize handlers
	dbHandler := handlers.NewDBHandler(dbService)

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/open", dbHandler.OpenDatabase)
		api.GET("/buckets", dbHandler.ListBuckets)
		api.POST("/bucket", dbHandler.CreateBucket)
		api.DELETE("/bucket/:name", dbHandler.DeleteBucket)
		api.GET("/bucket/:name/keys", dbHandler.ListKeys)
		api.GET("/bucket/:name/key/:key", dbHandler.GetValue)
		api.PUT("/bucket/:name/key/:key", dbHandler.SetValue)
		api.DELETE("/bucket/:name/key/:key", dbHandler.DeleteKey)
	}

	// Start server
	log.Println("Server starting on :8080")
	r.Run(":8080")
}
