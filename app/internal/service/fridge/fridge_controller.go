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
// Create is action: POST /fridges
func (ctrl *Controller) Create(c *gin.Context) {
	var service Service
	fridge,err := service.CreateModel(c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusCreated, fridge)
}
// ReadByID is action: GET /fridges?fid={fridge_id}
func (ctrl *Controller) ReadByID(c *gin.Context) {
	var service Service
	id := c.Query("fid")
	fridge,err := service.GetByID(id)
	// log.Println(fridge.Name)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, fridge)
	}
	
}
//Update action: PUT /users?fid={fridge_id}
func (ctrl *Controller) Update(c *gin.Context) {
	var service Service
	id := c.Query("fid")
	fridge,err := service.UpdateByID(id,c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusOK, fridge)
	}
	
}
// Delete action: DELETE /users?fid={fridge_id}
func (ctrl *Controller) Delete(c *gin.Context) {
	var service Service
	id := c.Query("fid")
	if err := service.DeleteByID(id); err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.JSON(http.StatusNoContent, gin.H{"id #"+ id: "deleted successfully"})
	}
	
}
// MyFridge action: GET /fridges/my-fridge?uid={user_id}
func (ctrl *Controller) GetMyFridge(c *gin.Context) {
	var service Service
	uid := c.Query("uid")
	fridge,err := service.MyFridge(uid)
	if err!=nil {
		log.Println(err)
		c.Abort()
	} else {
		c.JSON(http.StatusOK, fridge)
	}
}

func (ctrl *Controller) GetFollowFridge(c *gin.Context) {
	var service Service
	uid := c.Query("uid")
	flist,err := service.GetFollowFridgeList(uid)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, flist)
	}
}