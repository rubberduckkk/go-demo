package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(ResponseInterceptor, AnotherMiddleware)
	router.GET("/bind_bool", bindBool)
	router.POST("/post_form", postForm)

	router.Run("localhost:8080")
}

type BindBoolReq struct {
	IsOK bool `json:"is_ok" form:"is_ok"`
}

func bindBool(c *gin.Context) {
	req := new(BindBoolReq)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	RspWithResponse(c, 123, "ok", nil)
}

func postForm(c *gin.Context) {
	testKey := "test"
	c.JSON(http.StatusOK, gin.H{
		"key": testKey,
		"val": c.PostForm(testKey),
	})
}
