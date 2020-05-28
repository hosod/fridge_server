package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/service/user"
	"github.com/hosod/fridge_server/app/internal/service/fridge"
	"github.com/hosod/fridge_server/app/internal/service/food_genre"
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

	f := router.Group("/fridges")
	{
		fridgeCtrl := fridge.Controller{}
		f.GET("", fridgeCtrl.ReadAll)
		f.POST("", fridgeCtrl.Create)
		f.GET("/:id", fridgeCtrl.ReadByID)
		f.PUT("/:id", fridgeCtrl.Update)
		f.DELETE("/:id", fridgeCtrl.Delete)
		f.GET("/:id/users", fridgeCtrl.GetUserList)
	}

	fg := router.Group("/food_genres")
	{
		foodGenreCtrl := food_genre.Controller{}
		fg.GET("", foodGenreCtrl.ReadAll)
		fg.POST("", foodGenreCtrl.Create)
		fg.GET("/:id", foodGenreCtrl.ReadByID)
		fg.PUT("/:id", foodGenreCtrl.Update)
		fg.DELETE("/:id", foodGenreCtrl.Delete)
	}

	return router
}