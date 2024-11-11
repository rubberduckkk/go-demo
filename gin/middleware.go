package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var responseWriterPool = sync.Pool{
	New: func() interface{} {
		return &responseWriter{}
	},
}

type responseWriter struct {
	gin.ResponseWriter
	data *codeStruct
}

type codeStruct struct {
	Code int `json:"code"`
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.data = new(codeStruct)
	if err := json.Unmarshal(b, w.data); err != nil {
		w.data.Code = -1
	}
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) Reset() {
	w.ResponseWriter = nil
	w.data = nil
}

func ResponseInterceptor(c *gin.Context) {
	w := responseWriterPool.Get().(*responseWriter)
	w.ResponseWriter = c.Writer
	c.Writer = w

	c.Next()

	log.Printf("Path: %v Code: %v", c.Request.URL.Path, w.data.Code)

	c.Writer = w.ResponseWriter
	w.Reset()
	responseWriterPool.Put(w)
}

func AnotherMiddleware(c *gin.Context) {
	c.Next()
	code := c.Writer.Status()
	log.Printf("another middleware: code: %v", code)
}

// Response http接口的通用返回结构，一般不直接使用
func Response(code int, message string, data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = code
	result["message"] = message
	result["data"] = data
	return result
}

// RspWithResponse 最底层的通用返回
func RspWithResponse(c *gin.Context, code int, msg string, data interface{}) {
	dataType := c.Request.FormValue("data_type")
	if dataType == "" {
		dataType = c.Query("data_type")
	}
	c.JSON(http.StatusOK, Response(code, msg, data))
}
