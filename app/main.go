package main

import (
	// "log"
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
)

func main() {
	var dev = flag.Bool("dev", false, "please specify -dev flag")

	flag.Parse()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Wolrd hoge fuga piyo12345678\n")
	})
	// v1 := router.Group("api/v1")
	// {
	// 	v1.POST("/user", users.Register)
	// }
	database.Init(*dev)
	defer database.Close()
	router.Run(":9000")
}
