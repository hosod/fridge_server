package content

import (
	"log"
	"time"


	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)

type Service struct {}
type Content entity.Content
type ContentResult struct{
	ID int `json:"id"`
	Name string `json:"name"`
	ExpirationDate string `json:"expiration_date"`
	Quantity float32 `json:"quantity"`
	Genre entity.FoodGenre `json:"genre"`
}

var tokyo,_ = time.LoadLocation("Asia/Tokyo")

//PostForm is struct for contents-posting form
type PostForm struct {
	ExpirationDate string `json:"expiration_date"`
	Quantity float32 `json:"quantity"`
	FridgeID int `json:"fridge_id"`
	FoodTypeID int `json:"food_type_id"`
}
// GetByID ois read data from id
func (s *Service) GetByID(id string) (ContentResult, error) {
	db := database.GetDB()
	var contentResult ContentResult
	var content Content
	var foodType entity.FoodType
	var foodGenre entity.FoodGenre

	db.Where("id=?", id).First(&content)
	// log.Println(content.FoodType)
	db.Where("id=?",content.FoodTypeID).First(&foodType)
	db.Where("id=?", foodType.GenreID).First(&foodGenre)
	contentResult = ContentResult{
		ID:content.ID,
		Name:foodType.Name,
		ExpirationDate: content.ExpirationDate.Format("2006/01/02"),
		Quantity: content.Quantity,
		Genre: foodGenre,
	}

	return contentResult,nil
}
// CreateModel create content data and return it
func (s *Service)CreateModel(c *gin.Context) (Content,error) {
	db := database.GetDB()
	var contentForm PostForm
	var content Content
	if err:=c.BindJSON(&contentForm);err!=nil {
		log.Println("BindJSON error")
		return content,err
	}
	date,err := time.Parse("2006/01/02", contentForm.ExpirationDate)
	if err!=nil {
		log.Println(err)
		return content,err
	}
	content = Content{
		ExpirationDate:date,
		Quantity:contentForm.Quantity,
		FridgeID:contentForm.FridgeID,
		FoodTypeID:contentForm.FoodTypeID,
	}

	if err:=db.Create(&content).Error; err!=nil {
		return content,err
	}
	return content,nil
}
