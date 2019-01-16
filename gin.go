package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func gin1() {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.GET("/login", loginEndpoint)

	v2 := r.Group("/v1")
	
	fmt.Println(v2)
	r.Run(":80")
}

func test() {

}

func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "loginEndpoint")
}