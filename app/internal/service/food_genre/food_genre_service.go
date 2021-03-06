package food_genre

import (
	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)

// Service is related to food genre
type Service struct{}

// FoodGenre is alias of entity.FoodGenre struct
type FoodGenre entity.FoodGenre

// FoodElemt is info of a summary of food
type FoodElement struct {
	FoodTypeID int       `json:"food_type_id"`
	Name       string    `json:"name"`
	Image      string    `json:"img_url"`
	Genre      FoodGenre `json:"genre"`
}

type FoodList struct {
	Foods []*FoodElement `json:"food_list"`
}

// GetAll return list of all food genres
func (s *Service) GetAll() ([]FoodGenre, error) {
	db := database.GetDB()
	var food_genres []FoodGenre

	// SELECT * FROM food_genres;
	if err := db.Find(&food_genres).Error; err != nil {
		return nil, err
	}

	return food_genres, nil
}

// CreateModel create FoodGenre model and return it
func (s *Service) CreateModel(c *gin.Context) (FoodGenre, error) {
	db := database.GetDB()
	var food_genre FoodGenre

	if err := c.BindJSON(&food_genre); err != nil {
		return food_genre, err
	}
	// INSERT INTO food_genres values({ID},{NAME},{UNIT})
	if err := db.Create(&food_genre).Error; err != nil {
		return food_genre, err
	}
	return food_genre, nil
}

// GetByID is get a FoodGenre struct searching by id
func (s *Service) GetByID(id string) (FoodGenre, error) {
	db := database.GetDB()
	var food_genre FoodGenre
	// SELECT * FROM food_genres WHERE id={ID};
	if err := db.Where("id = ?", id).First(&food_genre).Error; err != nil {
		return food_genre, err
	}
	return food_genre, nil
}

// UpdateByID is update food genre
func (s *Service) UpdateByID(id string, c *gin.Context) (FoodGenre, error) {
	db := database.GetDB()
	var food_genre FoodGenre
	food_genre, err := s.GetByID(id)
	if err != nil {
		return food_genre, err
	}
	if err = c.BindJSON(&food_genre); err != nil {
		return food_genre, err
	}

	if err = db.Save(&food_genre).Error; err != nil {
		return food_genre, err
	}
	return food_genre, nil
}

// DeleteByID is delete a food genre
func (s *Service) DeleteByID(id string) error {
	db := database.GetDB()
	var food_genre FoodGenre

	if err := db.Where("id=?", id).Delete(&food_genre).Error; err != nil {
		return err
	}
	return nil
}

// GetImgByID returns icon image for image_id
func (s *Service) GetImgByID(id string, c *gin.Context) string {
	imgVege := "/go/src/github.com/hosod/fridge_server/app/imgs/flower_vegebouquet.png"
	imgMeat := "/go/src/github.com/hosod/fridge_server/app/imgs/food_niku_pack.png"
	imgVeba := "/go/src/github.com/hosod/fridge_server/app/imgs/soda6_skyblue.png"

	if id == "vege" {
		return imgVege
	} else if id == "meat" {
		return imgMeat
	} else {
		return imgVeba
	}
}

// GetWholeNameList returns food genre name list
func (s *Service) GetWholeNameList() (FoodList, error) {
	db := database.GetDB()

	foodList := FoodList{make([]*FoodElement, 0)}

	rows, err := db.Table("food_types").Select("*").
		Joins("JOIN food_genres on food_types.genre_id=food_genres.id").Rows()
	if err != nil {
		return foodList, err
	}
	for rows.Next() {
		var ftID, fgID, tmp int
		var ftName, img, fgName, fgUnit string
		rows.Scan(&ftID, &ftName, &img, &tmp, &fgID, &fgName, &fgUnit)
		element := FoodElement{
			FoodTypeID: ftID,
			Name:       ftName,
			Image:      img,
			Genre:      FoodGenre{ID: fgID, Name: fgName, Unit: fgUnit},
		}
		foodList.Foods = append(foodList.Foods, &element)
	}
	return foodList, nil

}
