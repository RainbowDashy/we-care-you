package main

import (
	"github.com/RainbowDashy/we-care-you/store"
	"github.com/gin-gonic/gin"
)

type API struct {
	g *gin.RouterGroup
	s *store.Store
}

func NewAPI(g *gin.RouterGroup, s *store.Store) *API {
	return &API{
		g: g,
		s: s,
	}
}

func (a *API) Register() {
	a.g.GET("/users", a.getUser)
	a.g.POST("/users", a.postUser)
}

// For authorization reason, not implemented yet
func (a *API) getUser(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func (a *API) postUser(c *gin.Context) {
	user := &store.User{}
	if err := c.ShouldBind(user); err != nil {
		c.JSON(400, gin.H{
			"msg": err,
		})
		return
	}
	if err := a.s.InsertUser(user); err != nil {
		c.JSON(500, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(200, user)
}
