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
	s, err := store.NewStore(config.dbPath)
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	api := NewAPI(r.Group("/api"), s)
	api.Register()

	r.Run()
}
