package fridge

import(
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
// Controller is 
type Controller struct{}

//ReadAll is  read
func (ctrl *Controller) ReadAll(c *gin.Context) {
	var service Service
	fridges,err := service.GetAll()
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK,fridges)
}
// Create is 
func (ctrl *Controller) Create(c *gin.Context) {
	var service Service
	fridge,err := service.CreateModel(c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusCreated, fridge)
}
// ReadByID is 
func (ctrl *Controller) ReadByID(c *gin.Context) {
	var service Service
	fridge,err := service.GetByID(c.Params.ByName("id"))
	log.Println(fridge.Name)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, fridge)
}
//Update action: PUT /users/:id
func (ctrl *Controller) Update(c *gin.Context) {
	var service Service
	id := c.Params.ByName("id")
	fridge,err := service.UpdateByID(id,c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, fridge)
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
// GetUserList is action: GET /users/:id/fridges
func (ctrl *Controller) GetUserList(c *gin.Context) {
	var service Service
	fridgeID := c.Params.ByName("id")
	users,err := service.GetUserList(fridgeID)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, users)
}