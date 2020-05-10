package user

import(
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

)
// Controller is user controller
type Controller struct{}

// Index action:   GET /users
func (ctrl *Controller) Index(c *gin.Context) {
	var service Service
	users,err := service.GetAll()
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, users)
}
// Create action: POST /users
func (ctrl *Controller) Create(c *gin.Context) {
	var service Service
	user,err := service.CreateModel(c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusCreated, user)
}
// ShowByID is show user idで指定されたやつ
func (ctrl *Controller) ShowByID(c *gin.Context) {
	var service Service
	user,err := service.GetByID(c.Params.ByName("id"))
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}
