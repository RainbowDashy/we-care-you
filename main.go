package main

import (
	"github.com/RainbowDashy/we-care-you/store"
	"github.com/gin-gonic/gin"
)

type Config struct {
	dbPath string
}

var config Config

func init() {
	config = Config{
		dbPath: "./data.db",
	}
}

func main() {
	_, err := store.NewStore(config.dbPath)
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
