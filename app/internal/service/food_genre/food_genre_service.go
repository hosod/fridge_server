package food_genre

import(
	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"

    "os"
)

// Service is related to food genre
type Service struct{}
// FoodGenre is alias of entity.FoodGenre struct
type FoodGenre entity.FoodGenre

// GetAll return list of all food genres
func (s *Service) GetAll() ([]FoodGenre, error) {
	db := database.GetDB()
	var food_genres []FoodGenre

	// SELECT * FROM food_genres;
	if err:=db.Find(&food_genres).Error; err!=nil {
		return nil,err
	}

	return food_genres, nil
}
// CreateModel create FoodGenre model and return it
func (s *Service) CreateModel(c *gin.Context) (FoodGenre, error) {
	db := database.GetDB()
	var food_genre FoodGenre

	if err:=c.BindJSON(&food_genre); err!=nil {
		return food_genre,err
	}
	// INSERT INTO food_genres values({ID},{NAME},{UNIT})
	if err:=db.Create(&food_genre).Error; err!=nil {
		return food_genre,err
	}
	return food_genre,nil
}
// GetByID is get a FoodGenre struct searching by id
func (s *Service) GetByID(id string) (FoodGenre, error) {
	db := database.GetDB()
	var food_genre FoodGenre
// SELECT * FROM food_genres WHERE id={ID};
	if err := db.Where("id = ?", id).First(&food_genre).Error; err!=nil {
		return food_genre,err
	}
	return food_genre,nil
}
// UpdateByID is update food genre
func (s *Service) UpdateByID(id string, c *gin.Context) (FoodGenre,error) {
	db := database.GetDB()
	var food_genre FoodGenre
	food_genre,err := s.GetByID(id)
	if err!=nil {
		return food_genre,err
	}
	if err=c.BindJSON(&food_genre); err!=nil {
		return food_genre,err
	}

	if err=db.Save(&food_genre).Error; err!=nil {
		return food_genre,err
	}
	return food_genre,nil
}
// DeleteByID is delete a food genre
func (s *Service) DeleteByID(id string) error {
	db := database.GetDB()
	var food_genre FoodGenre

	if err := db.Where("id=?",id).Delete(&food_genre).Error; err!=nil {
		return err
	}
	return nil
}

func (s * Service) GetImgByID(id string) (*File, error) {
	img_vege, _ := os.Open("../../imgs/flower_vegebouquet.png")
	img_meat, _ := os.Open("../../imgs/food_niku_pack.png")
	img_veba, _ := os.Open("../../imgs/soda6_skyblue.png")
	if id == "vege" {
		return img_vege, nil
	} else if id == "meat" {
		return img_meat, nil
	} else if id == "veba" {
		return img_veba, nil
	} else {
		return img_vege, nil
	}
}

// GetWholeNameList returns food genre name list
func (s * Service) GetWholeNameList() ([]string, error){
	db := database.GetDB()
	var names []string
	db.Table("food_genres").Pluck("name", &names)

	return names,nil
}


