package main

import (
	"encoding/json"
	"log"
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
	code int
}

type codeStruct struct {
	Code int `json:"code"`
}

func (w *responseWriter) Write(b []byte) (int, error) {
	var c codeStruct
	if err := json.Unmarshal(b, &c); err == nil {
		w.code = c.Code
	} else {
		w.code = -1
	}
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) Reset() {
	w.code = 0
}

func ResponseInterceptor(c *gin.Context) {
	w := responseWriterPool.Get().(*responseWriter)
	w.ResponseWriter = c.Writer
	c.Writer = w

	c.Next()

	log.Println("Response Code:", w.code)

	w.Reset()
	responseWriterPool.Put(w)
}
