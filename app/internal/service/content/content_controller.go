package content

import (
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
	contentResult,err := service.GetByID(contentID)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contentResult)
	}
}