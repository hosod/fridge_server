package content

import (
	"log"
	"time"


	// "github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)

type Service struct {}
type Content entity.Content
type ContentResult struct{
	ID int `json:"id"`
	Name string `json:"name"`
	ExpirationDate time.Time `json:"expiration_date"`
	Quantity float32 `json:"quantity"`
	Genre entity.FoodGenre `json:"genre"`
}

func (s *Service) GetByID(id string) (ContentResult, error) {
	db := database.GetDB()
	var contentResult ContentResult
	var content Content
	var foodType entity.FoodType
	var foodGenre entity.FoodGenre

	db.Where("id=?", id).First(&content)
	// log.Println(content.FoodType)
	db.Where("id=?",content.FoodTypeID).First(foodType)
	db.Where("id=?", foodType.GenreID).First(foodGenre)
	contentResult = ContentResult{
		ID:content.ID,
		Name:foodType.Name,
		ExpirationDate: content.ExpirationDate,
		Quantity: content.Quantity,
		Genre: foodGenre,
	}

	return contentResult,nil


}