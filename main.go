package main

import (
	"database/sql"

	"github.com/RainbowDashy/we-care-you/store"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func init() {
	db, _ = store.OpenDatabase()
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
