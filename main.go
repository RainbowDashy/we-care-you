package main

import (
	"github.com/RainbowDashy/we-care-you/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	s, err := store.NewStore(config.dbPath)
	l, _ := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	// cors
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Header("Access-Control-Allow-Methods", "POST, GET, PATCH")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := NewAPI(r.Group("/api"), s, l)
	api.Register()

	r.Run()
}
