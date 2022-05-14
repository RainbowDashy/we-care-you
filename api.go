package main

import (
	"github.com/RainbowDashy/we-care-you/store"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type API struct {
	g *gin.RouterGroup
	s *store.Store
}

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func NewAPI(g *gin.RouterGroup, s *store.Store) *API {
	return &API{
		g: g,
		s: s,
	}
}

func (a *API) NewAuthMiddleware() *jwt.GinJWTMiddleware {
	m, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "we-care-you",
		// Should be random string
		// Use "test" in development environment
		Key: []byte("test"),
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*store.User); ok {
				user, _ := a.s.GetUserByUsername(v.Username)
				return jwt.MapClaims{
					"id":       user.Id,
					"username": user.Username,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			input := &User{}
			if err := c.ShouldBind(input); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			user := store.NewUser(input.Username, input.Password)
			if a.s.ValidUser(user) {
				return user, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
	})
	return m
}

func (a *API) Register() {
	authMiddleware := a.NewAuthMiddleware()
	a.g.POST("/login", authMiddleware.LoginHandler)
	a.g.POST("/logout", authMiddleware.LogoutHandler)
	a.g.GET("/users", authMiddleware.MiddlewareFunc(), a.getUser)
	a.g.POST("/users", a.postUser)
}

func (a *API) getUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"id":       claims["id"],
		"username": claims["username"],
	})
}

func (a *API) postUser(c *gin.Context) {
	input := &User{}
	if err := c.ShouldBind(input); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	user := store.NewUser(input.Username, input.Password)
	if err := a.s.InsertUser(user); err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"id": user.Id,
	})
}
