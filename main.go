package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error reading env variables: " + err.Error())
	}

	r := setUpRouter()
	err = r.Run()
	if err != nil {
		panic("Server start failed: " + err.Error())
	}
}

func setUpRouter() *gin.Engine {
	app := gin.Default()

	app.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hey!",
		})
	})

	return app
}