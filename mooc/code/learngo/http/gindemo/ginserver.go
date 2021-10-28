package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// middleware
	const keyRequestId = "requestId"
	r.Use(func(c *gin.Context) {
		s := time.Now()

		c.Next()

		// 不管访问什么,都能先进到这里面来
		// log latency, response code
		logger.Info("incoming request:",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
		//log.Fatalf(c.Request.URL.Path)
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {

		h := gin.H{
			"message": "pong",
		}

		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello gin")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
