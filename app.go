package main

import (
	"example.com/community-app-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("failed to load .env file")
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")

}
