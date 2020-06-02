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
// ReadByID action: GET /users?uid={user_id}
func (ctrl *Controller) ReadByID(c *gin.Context) {
	var service Service
	userID := c.Query("uid")
	user,err := service.GetByID(userID)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}

// Update action: PUT /users?uid={user_id}
func (ctrl *Controller) Update(c *gin.Context) {
	var service Service
	userID := c.Query("uid")
	user,err := service.UpdateByID(userID, c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, user)
}
// Delete action: DELETE /users?uid={user_id}
func (ctrl *Controller) Delete(c *gin.Context) {
	var service Service
	userID := c.Query("uid")
	if err := service.DeleteByID(userID); err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusForbidden)
	}
	c.JSON(http.StatusNoContent, gin.H{"id #"+ userID: "deleted successfully"})
}




