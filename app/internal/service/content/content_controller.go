package content

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
// Controller is
type Controller struct{}

// ReadByID is action: GET /contents?cid={content_id}
func (ctrl *Controller) ReadByID(c *gin.Context) {
	var service Service
	contentID := c.Query("cid")

	hoge := c.Query("hoge")
	log.Println("hoge:",hoge)
	contentResult,err := service.GetByID(contentID)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contentResult)
	}
}
// Create is action: POST /contents
func (ctrl *Controller) Create(c *gin.Context) {
	var service Service
	contentList,err := service.CreateModel(c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		c.JSON(http.StatusOK, gin.H{"message":fmt.Sprintf("%d contents are registered successfully", len(contentList))})
	}
}
// Delete is action: DELETE /contents?cid={content_id}
func (ctrl *Controller) Delete(c *gin.Context) {
	var service Service
	id := c.DefaultQuery("cid", "0")
	err := service.DeleteByID(id)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusNoContent, gin.H{"id #"+ id: "deleted successfully"})
	}
}
// ReadByFridgeID is action: GET /contents/fridge?fid={fridge_id}
func (ctrl *Controller) ReadByFridgeID(c *gin.Context) {
	var service Service
	fridgeID := c.Query("fid")
	contentResultList,err := service.GetByFridgeID(fridgeID)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contentResultList)
	}
}
// ReadByUserID is action: GET /contents/user?uid={user_id}
func (ctrl *Controller) ReadByUserID(c * gin.Context) {
	var service Service
	userID := c.Query("uid")
	contentResultList,err := service.GetByUserID(userID)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contentResultList)
	}
}