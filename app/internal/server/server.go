package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/service"
)

// Init is initialize server
func Init() {
	router := makeRouter()
	router.Run(":9000")
}

func makeRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world.\n")
	})

	u := router.Group("/users") 
	{
		userCtrl := user.Controller{}
		u.GET("", userCtrl.ReadAll)
		u.POST("", userCtrl.Create)
		u.GET("/:id", userCtrl.ReadByID)
		u.PUT("/:id", userCtrl.Update)
		u.DELETE("/:id", userCtrl.Delete)
	}

	return router
}