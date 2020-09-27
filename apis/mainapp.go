package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()

	var configdata ConfigData = readConfig("config.json")
	fmt.Printf("hello, లక్ష్మణ \t" + configdata.ServerPort)

	ConnectDB()

	router.GET("/api/login/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("/api/register/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run(":" + configdata.ServerPort)
}
