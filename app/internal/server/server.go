package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/service/user"
	"github.com/hosod/fridge_server/app/internal/service/fridge"
	"github.com/hosod/fridge_server/app/internal/service/food_genre"
	"github.com/hosod/fridge_server/app/internal/service/content"
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
		// u.GET("", userCtrl.ReadAll)
		u.POST("", userCtrl.Create)
		u.GET("", userCtrl.ReadByID)
		u.PUT("", userCtrl.Update)
		u.DELETE("", userCtrl.Delete)
		
	}

	f := router.Group("/fridges")
	{
		fridgeCtrl := fridge.Controller{}
		// f.GET("", fridgeCtrl.ReadAll)
		f.GET("", fridgeCtrl.ReadByID)
		f.POST("", fridgeCtrl.Create)
		f.PUT("", fridgeCtrl.Update)
		f.DELETE("", fridgeCtrl.Delete)
		f.GET("/my-fridge", fridgeCtrl.GetMyFridge)
		f.GET("/follow-fridges", fridgeCtrl.GetFollowFridge)
		// f.GET("/:id/users", fridgeCtrl.GetUserList)
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

	c := router.Group("/contents")
	{
		contentCtrl := content.Controller{}
		c.GET("", contentCtrl.ReadByID)

	}

	return router
}