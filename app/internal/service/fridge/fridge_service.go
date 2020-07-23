package fridge

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)

//Service is related to fridge's behavior
type Service struct{}

// Fridge is alias of entity.Fridge struct
type Fridge entity.Fridge

type FridgeAndOwner struct {
	Fridge entity.Fridge `json:"fridge"`
	Owner  entity.User   `json:"user"`
}

type FollowFridgeList struct {
	FridgeAndOwner []*FridgeAndOwner `json:"pairs"`
}

// GetAll return list of all fridges
func (s *Service) GetAll() ([]Fridge, error) {
	db := database.GetDB()
	var fridges []Fridge

	if err := db.Find(&fridges).Error; err != nil {
		return nil, err
	}
	return fridges, nil
}

// CreateModel create Fridge model and return it
func (s *Service) CreateModel(c *gin.Context) (Fridge, error) {
	db := database.GetDB()
	var fridge Fridge

	if err := c.BindJSON(&fridge); err != nil {
		return fridge, err
	}
	if err := db.Create(&fridge).Error; err != nil {
		return fridge, err
	}
	return fridge, nil
}

// GetByID is get a Fridge struct searching by id
func (s *Service) GetByID(id string) (Fridge, error) {
	db := database.GetDB()
	var fridge Fridge

	if err := db.Where("id=?", id).First(&fridge).Error; err != nil {
		return fridge, err
	}
	return fridge, nil
}

// UpdateByID is update fridge
func (s *Service) UpdateByID(id string, c *gin.Context) (Fridge, error) {
	db := database.GetDB()
	var fridge Fridge
	fridge, err := s.GetByID(id)
	if err != nil {
		return fridge, err
	}

	if err = c.BindJSON(&fridge); err != nil {
		return fridge, err
	}

	if err = db.Save(&fridge).Error; err != nil {
		return fridge, err
	}
	return fridge, nil
}

// DeleteByID is delete a fridge
func (s *Service) DeleteByID(id string) error {
	db := database.GetDB()
	var fridge Fridge

	if err := db.Where("id=?", id).Delete(&fridge).Error; err != nil {
		return err
	}
	return nil
}

// GetUserList is return user list fridge have
func (s *Service) GetUserList(fridgeID string) ([]*entity.User, error) {
	db := database.GetDB()
	var fridge Fridge
	if err := db.Where("id=?", fridgeID).Preload("User").First(&fridge).Error; err != nil {
		return nil, err
	}
	return fridge.User, nil
}

// MyFridge return fridge owned by user(userID)
func (s *Service) MyFridge(userID string) (Fridge, error) {
	db := database.GetDB()
	var fridge Fridge
	uid, err := strconv.Atoi(userID)
	if err != nil {
		return fridge, err
	}
	if err = db.Joins("JOIN users ON users.my_fridge_id=fridges.id AND users.id=?", uid).Find(&fridge).Error; err != nil {
		return fridge, err
	}
	return fridge, nil
}

func (s *Service) GetFollowFridgeList(userID string) (FollowFridgeList, error) {
	db := database.GetDB()
	var followFridgeList FollowFridgeList

	rows, err := db.Table("fridges").
		Select("fridges.id,fridges.name,users.id,users.name,users.email").
		Joins("JOIN user_follow_fridge as follow ON fridges.id=follow.fridge_id AND follow.user_id=?", userID).
		Joins("LEFT OUTER JOIN users ON fridges.id = users.my_fridge_id").Rows()
	if err != nil {
		return followFridgeList, err
	}

	followFridgeList.FridgeAndOwner = make([]*FridgeAndOwner, 0)
	for rows.Next() {
		var uid, fid int
		var uname, uemail, fname string
		rows.Scan(&fid, &fname, &uid, &uname, &uemail)
		log.Println(uid, uname, uemail, fid, fname)
		fridgeAndOwner := FridgeAndOwner{
			Owner:  entity.User{ID: uid, Name: uname, Email: uemail},
			Fridge: entity.Fridge{ID: fid, Name: fname},
		}
		followFridgeList.FridgeAndOwner = append(followFridgeList.FridgeAndOwner, &fridgeAndOwner)
	}
	// log.Println(rows)
	return followFridgeList, nil

}
