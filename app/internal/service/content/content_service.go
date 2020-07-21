package content

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)

type Service struct{}

// Content is alias of entity.Content
type Content entity.Content

// ContentResult is response data format for food data
type ContentResult struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	ExpirationDate string           `json:"expiration_date"`
	Quantity       float32          `json:"quantity"`
	Image          string           `json:"image"`
	Genre          entity.FoodGenre `json:"genre"`
}

// ContentResultList is reponse data format for
type ContentResultList struct {
	Foods []*ContentResult `json:"foods"`
}

var tokyo, _ = time.LoadLocation("Asia/Tokyo")

//PostForm is struct for contents-posting form
type PostForm struct {
	ExpirationDate string  `json:"expiration_date"`
	Quantity       float32 `json:"quantity"`
	FridgeID       int     `json:"fridge_id"`
	UserID         int     `json:"user_id"`
	FoodTypeID     int     `json:"food_type_id"`
}

// PostFormList is data format for endpoint POST: /contents
type PostFormList struct {
	FormList []*PostForm `json:"foods"`
}

// GetByID ois read data from id
func (s *Service) GetByID(id string) (ContentResult, error) {
	db := database.GetDB()
	var contentResult ContentResult
	var content Content
	var foodType entity.FoodType
	var foodGenre entity.FoodGenre

	if err := db.Where("id=?", id).First(&content).Error; err != nil {
		return contentResult, err
	}
	if err := db.Where("id=?", content.FoodTypeID).First(&foodType).Error; err != nil {
		return contentResult, err
	}
	if err := db.Where("id=?", foodType.GenreID).First(&foodGenre).Error; err != nil {
		return contentResult, err
	}
	contentResult = ContentResult{
		ID:             content.ID,
		Name:           foodType.Name,
		ExpirationDate: content.ExpirationDate.Format("2006/01/02"),
		Quantity:       content.Quantity,
		Image:          foodType.Image,
		Genre:          foodGenre,
	}

	return contentResult, nil
}

// UpdateByID update quantity value of content
func (s *Service) UpdateByID(id string, c *gin.Context) error {
	db := database.GetDB()
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	content := entity.Content{ID: i}

	var postForm PostForm
	if err := c.BindJSON(&postForm); err != nil {
		return err
	}

	err = db.Model(&content).Updates(
		map[string]interface{}{
			"quantity": postForm.Quantity,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID is delete a content
func (s *Service) DeleteByID(id string) error {
	db := database.GetDB()
	var content Content
	if err := db.Where("id=?", id).Delete(&content).Error; err != nil {
		return err
	}
	return nil
}

// GetByFridgeID is return list of contents which are contained a fridge(friodgeID)
func (s *Service) GetByFridgeID(fridgeID string) (ContentResultList, error) {
	db := database.GetDB()
	var contentResultList ContentResultList

	rows, err := db.Table("contents").
		Select("contents.id,food_types.name,contents.quantity,contents.expiration_date,food_types.image,food_genres.id,food_genres.name,food_genres.unit").
		Joins("join food_types on food_types.id=contents.food_type_id and contents.fridge_id=?", fridgeID).
		Joins("join food_genres on food_genres.id=food_types.genre_id").Rows()
	if err != nil {
		return contentResultList, err
	}
	contentResultList.Foods = make([]*ContentResult, 0)
	for rows.Next() {
		var contentID, genreID int
		var typeName, contentsDate, image, genreName, genreUnit string
		var quantity float32
		rows.Scan(&contentID, &typeName, &quantity, &contentsDate, &image, &genreID, &genreName, &genreUnit)
		contentResult := ContentResult{
			ID:             contentID,
			Name:           typeName,
			ExpirationDate: contentsDate,
			Quantity:       quantity,
			Image:          image,
			Genre:          entity.FoodGenre{ID: genreID, Name: genreName, Unit: genreUnit},
		}
		contentResultList.Foods = append(contentResultList.Foods, &contentResult)
	}
	return contentResultList, nil
}

// GetByUserID is get list of contents which are owned by a user(userID)
func (s *Service) GetByUserID(userID string) (ContentResultList, error) {
	db := database.GetDB()
	var contentResultList ContentResultList
	var user entity.User
	if err := db.Where("id=?", userID).First(&user).Error; err != nil {
		return contentResultList, err
	}
	return s.GetByFridgeID(strconv.Itoa(user.MyFridgeID))
}

// CreateModel create content data and return it
func (s *Service) CreateModel(c *gin.Context) ([]Content, error) {
	db := database.GetDB()
	var postFormList PostFormList
	var content Content
	contentList := []Content{}
	if err := c.BindJSON(&postFormList); err != nil {
		log.Println("BindJSON error")
		return contentList, err
	}

	for _, postForm := range postFormList.FormList {
		var user entity.User
		if err := db.Where("id=?", postForm.UserID).First(&user).Error; err != nil {
			log.Println(err)
			continue
		}
		fridgeID := user.MyFridgeID

		date, err := time.Parse("2006/01/02", postForm.ExpirationDate)
		if err != nil {
			log.Println(err)
			continue
		}
		content = Content{
			ExpirationDate: date,
			Quantity:       postForm.Quantity,
			FridgeID:       fridgeID,
			FoodTypeID:     postForm.FoodTypeID,
		}
		if err := db.Create(&content).Error; err != nil {
			log.Println(err)
			continue
		}
		contentList = append(contentList, content)
	}

	return contentList, nil
}
