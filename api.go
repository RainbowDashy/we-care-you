package main

import (
	"strconv"

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
	Location string `json:"location" form:"location"`
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
			user := store.NewUser(input.Username, input.Password, input.Location)
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
	a.g.PATCH("/malls/:id", authMiddlewareFunc, a.patchMall)
	a.g.GET("/items", authMiddlewareFunc, a.getItems)
	a.g.GET("/items/:id", authMiddlewareFunc, a.getItemById)
	a.g.GET("/orders", authMiddlewareFunc, a.getOrders)
	a.g.POST("/orders", authMiddlewareFunc, a.postOrders)
}

func handleErr(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func (a *API) getUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, err := a.s.GetUserByUsername(claims["usrename"].(string))
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"id":       user.Id,
		"username": user.Username,
		"location": user.Location,
	})
}

func (a *API) postUser(c *gin.Context) {
	input := &User{}
	if err := c.ShouldBind(input); err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	user := store.NewUser(input.Username, input.Password, input.Location)
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
		BeginTime int64         `json:"begintime"`
		EndTime   int64         `json:"endtime"`
		State     int64         `json:"state"`
		Items     []*store.Item `json:"items"`
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
	mall := &store.Mall{
		UserId:    user.Id,
		BeginTime: input.BeginTime,
		EndTime:   input.EndTime,
		State:     input.State,
	}
	err = a.s.CreateMall(user, mall, input.Items)
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, mall)
}

func (a *API) patchMall(c *gin.Context) {
	input := &struct {
		BeginTime int64 `json:"begintime"`
		EndTime   int64 `json:"endtime"`
		State     int64 `json:"state"`
	}{}
	mallId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	mall, err := a.s.GetMallById(mallId)
	if err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	claims := jwt.ExtractClaims(c)
	userName := claims["username"].(string)
	user, err := a.s.GetUserByUsername(userName)
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	if mall.UserId != user.Id {
		handleErr(c, 400, "Unauthorized")
		return
	}
	mall.BeginTime = input.BeginTime
	mall.EndTime = input.EndTime
	mall.State = input.State
	if err := a.s.UpdateMall(mall); err != nil {
		handleErr(c, 500, err.Error())
		return
	}
}

func (a *API) getItems(c *gin.Context) {
	input := &struct {
		MallId int64 `json:"mallid" form:"mallid"`
	}{}
	if err := c.ShouldBind(input); err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	if input.MallId == 0 {
		handleErr(c, 400, "mallid is 0")
		return
	}
	items, err := a.s.GetItemsByMallId(input.MallId)
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, items)
}

func (a *API) getItemById(c *gin.Context) {
	itemId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	item, err := a.s.GetItemById(itemId)
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, item)
}

func (a *API) getOrders(c *gin.Context) {
	input := &struct {
		MallId int64 `json:"mallid" form:"mallid"`
		UserId int64 `json:"userid" form:"userid"`
		ItemId int64 `json:"itemid" form:"itemid"`
	}{}
	if err := c.ShouldBind(input); err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	if input.MallId == 0 && input.UserId == 0 && input.ItemId == 0 {
		handleErr(c, 400, "id equals to 0")
		return
	}
	var orders []*store.MallCustomer
	var err error
	if input.MallId != 0 {
		orders, err = a.s.GetOrdersByMallId(input.MallId)
	} else if input.UserId != 0 {
		orders, err = a.s.GetOrdersByUserId(input.UserId)
	} else if input.ItemId != 0 {
		orders, err = a.s.GetOrdersByItemId(input.ItemId)
	}
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, orders)
}

func (a *API) postOrders(c *gin.Context) {
	input := &struct {
		Orders []*store.MallCustomer `json:"orders"`
	}{}
	if err := c.ShouldBind(input); err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	if len(input.Orders) == 0 {
		handleErr(c, 400, "orders is nil")
		return
	}
	claims := jwt.ExtractClaims(c)
	user, err := a.s.GetUserByUsername(claims["username"].(string))
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	err = a.s.Buy(user, input.Orders)
	if err != nil {
		handleErr(c, 500, err.Error())
		return
	}
	c.JSON(200, gin.H{})
}
