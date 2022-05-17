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

	api := NewAPI(r.Group("/api"), s, l)
	api.Register()

	r.Run()
}
