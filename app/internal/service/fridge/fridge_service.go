package fridge

import(
	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)
//Service is related to fridge's behavior
type Service struct{}
// Fridge is alias of entity.Fridge struct
type Fridge entity.Fridge

// GetAll return list of all fridges
func (s *Service) GetAll() ([]Fridge, error) {
	db := database.GetDB()
	var fridges []Fridge

	if err := db.Find(&fridges).Error; err!=nil {
		return nil,err
	}
	return fridges,nil
}
// CreateModel create Fridge model and return it
func (s *Service) CreateModel(c *gin.Context) (Fridge,error) {
	db := database.GetDB()
	var fridge Fridge

	if err := c.BindJSON(&fridge); err!=nil {
		return fridge,err
	}
	if err := db.Create(&fridge).Error; err!=nil {
		return fridge,err
	}
	return fridge,nil
}
// GetByID is get a Fridge struct searching by id
func (s *Service) GetByID(id string) (Fridge, error) {
	db := database.GetDB()
	var fridge Fridge

	if err := db.Where("id=?", id).First(&fridge).Error; err!=nil {
		return fridge,err
	}
	return fridge,nil
}
// UpdateByID is update fridge
func (s *Service) UpdateByID(id string, c *gin.Context) (Fridge,error) {
	db := database.GetDB()
	var fridge Fridge
	fridge,err := s.GetByID(id)
	if err!=nil {
		return fridge,err
	}

	if err = c.BindJSON(&fridge); err!=nil {
		return fridge,err
	}

	if err = db.Save(&fridge).Error; err!=nil {
		return fridge,err
	}
	return fridge,nil
}
// DeleteByID is delete a fridge
func (s *Service) DeleteByID(id string) error {
	db := database.GetDB()
	var fridge Fridge

	if err := db.Where("id=?",id).Delete(&fridge).Error; err!=nil {
		return err
	}
	return nil
}

