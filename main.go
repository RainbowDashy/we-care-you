package main

import (
	"database/sql"

	"github.com/RainbowDashy/we-care-you/store"
	"github.com/gin-gonic/gin"
)

type Config struct {
	DBPath string
}

type Server struct {
	DB     *sql.DB
	Config Config
}

var server *Server

func init() {
	config := Config{
		DBPath: "./data.db",
	}
	db, err := store.OpenDatabase(config.DBPath)
	if err != nil {
		panic(err)
	}
	server = &Server{
		DB:     db,
		Config: config,
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
