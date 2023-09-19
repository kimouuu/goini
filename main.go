package main

import (
	"log"
    "github.com/gin-gonic/gin"
        
     
)
func RunConnectMongoDB() {
	_, err := database.ConnectMongoDB()
	log.Printf("Connecting to MongoDB...")
	if err != nil {
		log.Printf("Error when connecting to MongoDB: %s", err.Error())
	} else {
		log.Println("Successfully connected to MongoDB!")
	}
}

func RunConnectGinGonic() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Mila!",
		})
	})

	err := r.Run(":" + configs.EnvPort()) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Printf("Error when running server: %s", err.Error())
	}
}