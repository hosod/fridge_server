package user

import(
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
// Controller is user controller
type Controller struct{}

// ReadAll action: GET /users
func (ctrl *Controller) ReadAll(c *gin.Context) {
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
// ReadByID action: GET /users/:id
func (ctrl *Controller) ReadByID(c *gin.Context) {
	var service Service
	// id := c.Query("uid")
	user,err := service.GetByID(c.Params.ByName("id"))
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}

// Update action: PUT /users/:id
func (ctrl *Controller) Update(c *gin.Context) {
	var service Service
	user,err := service.UpdateByID(c.Params.ByName("id"), c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, user)
}
// Delete action: DELETE /users/:id
func (ctrl *Controller) Delete(c *gin.Context) {
	var service Service
	id := c.Params.ByName("id")
	if err := service.DeleteByID(id); err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusForbidden)
	}
	c.JSON(http.StatusNoContent, gin.H{"id #"+ id: "deleted successfully"})
}




