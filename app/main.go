package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Wolrd hoge fuga piyo\n")
	})

	v1 := router.Group("api/v1")
	{
		v1.POST("/user", users.Register)
	}

	router.Run(":9000")
}
