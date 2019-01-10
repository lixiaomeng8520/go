package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func gin1() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "/user")
	})

	r.GET("/user/", func(c *gin.Context) {
		c.String(http.StatusOK, "/user/")
	})

	r.Run()
}