package food_genre

import(
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
// Controller is food genre controller
type Controller struct{}

// ReadAll action: GET /food_genres
func (ctrl *Controller) ReadAll(c *gin.Context) {
	var service Service
	food_genres,err := service.GetAll()
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, food_genres)
}
// Create action: POST /food_genres
func (ctrl *Controller) Create(c *gin.Context) {
	var service Service
	food_genre,err := service.CreateModel(c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusCreated, food_genre)
}
// ReadByID action: GET /food_genres/:id
func (ctrl *Controller) ReadByID(c *gin.Context) {
	var service Service
	food_genre,err := service.GetByID(c.Query("id"))
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, food_genre)
}

// Update action: PUT /food_genres/:id
func (ctrl *Controller) Update(c *gin.Context) {
	var service Service
	food_genre,err := service.UpdateByID(c.Query("id"), c)
	if err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, food_genre)
}
// Delete action: DELETE /food_genres/:id
func (ctrl *Controller) Delete(c *gin.Context) {
	var service Service
	id := c.Query("id")
	if err := service.DeleteByID(id); err!=nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusForbidden)
	}
	c.JSON(http.StatusNoContent, gin.H{"id #"+ id: "deleted successfully"})
}

// ReadImgByID action: GET /food_genres/:iid
func (ctrl *Controller) ReadImgByID(c *gin.Context) {
	var service Service
	img_food_genre := service.GetImgByID(c.Query("iid"), c)
	// if err!=nil {
	//	log.Println(err)
	//	c.AbortWithStatus(http.StatusNotFound)
	// }
	c.File(img_food_genre)
}


