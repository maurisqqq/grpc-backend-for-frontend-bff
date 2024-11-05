package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	data, err := doHello(c, "Maulana")
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, data)
	return
}

func helloRoute(router *gin.RouterGroup) {
	router.GET("/", hello)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	helloRoute(api.Group("/hello"))

	r.Run()
}
