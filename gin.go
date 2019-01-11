package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func gin1() {
	r := gin.Default()

	r.GET("/user/age", userAge)

	r.GET("/user/name", userName)

	r.GET("/book", book)

	r.Run()
}

func test() {

}

func userAge(c *gin.Context) {
	c.String(http.StatusOK, "/user/age")
}

func userName(c *gin.Context) {
	c.String(http.StatusOK, "/user/name")
}

func book(c *gin.Context) {
	c.String(http.StatusOK, "/book")
}