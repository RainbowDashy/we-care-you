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
	authMiddlewareFunc := authMiddleware.MiddlewareFunc()
	a.g.POST("/login", authMiddleware.LoginHandler)
	a.g.POST("/logout", authMiddleware.LogoutHandler)
	a.g.GET("/users", authMiddlewareFunc, a.getUser)
	a.g.POST("/users", a.postUser)
	a.g.GET("/malls", authMiddlewareFunc, a.getMalls)
	a.g.POST("/malls", authMiddlewareFunc, a.postMalls)
}

func handleErr(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
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
		handleErr(c, 400, err.Error())
		return
	}
	user := store.NewUser(input.Username, input.Password)
	if err := a.s.InsertUser(user); err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"id": user.Id,
	})
}

func (a *API) getMalls(c *gin.Context) {
	input := &struct {
		UserId int64 `json:"userid" form:"userid"`
	}{}
	if err := c.ShouldBind(input); err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	var malls []*store.Mall
	var err error
	if input.UserId == 0 {
		malls, err = a.s.GetMalls()
	} else {
		malls, err = a.s.GetMallsByUserId(input.UserId)
	}
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, malls)
}

func (a *API) postMalls(c *gin.Context) {
	input := &struct {
		Items []*store.Item `json:"items"`
	}{}
	if err := c.ShouldBind(input); err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	if len(input.Items) == 0 {
		handleErr(c, 400, "items is nil")
		return
	}
	claims := jwt.ExtractClaims(c)
	user, err := a.s.GetUserByUsername(claims["username"].(string))
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	mall, err := a.s.CreateMall(user, input.Items)
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, mall)
}
