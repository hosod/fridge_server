package user

import(
	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)

// Service is related to user's behavior
type Service struct{}
// User is alias of entity.User struct
type User entity.User



// GetAll return list of all users
func (s *Service) GetAll() ([]User, error) {
	db := database.GetDB()
	var users []User

	// SELECT * FROM users;
	if err:=db.Find(&users).Error; err!=nil {
		return nil,err
	}

	return users, nil
}
// CreateModel create User model and return it
func (s *Service) CreateModel(c *gin.Context) (User, error) {
	db := database.GetDB()
	var user User

	if err:=c.BindJSON(&user); err!=nil {
		return user,err
	}
	// INSERT INTO users values({ID},{NAME},{EMAIL})
	if err:=db.Create(&user).Error; err!=nil {
		return user,err
	}
	return user,nil
}
// GetByID is get a User struct searching by id
func (s *Service) GetByID(id string) (User, error) {
	db := database.GetDB()
	var user User
// SELECT * FROM users WHERE id={ID};
	if err := db.Where("id = ?", id).First(&user).Error; err!=nil {
		return user,err
	}
	return user,nil
}
// UpdateByID is update user
func (s *Service) UpdateByID(id string, c *gin.Context) (User,error) {
	db := database.GetDB()
	var user User
	user,err := s.GetByID(id)
	if err!=nil {
		return user,err
	}
	if err=c.BindJSON(&user); err!=nil {
		return user,err
	}

	if err=db.Save(&user).Error; err!=nil {
		return user,err
	}
	return user,nil
}
// DeleteByID is delete a user
func (s *Service) DeleteByID(id string) error {
	db := database.GetDB()
	var user User

	if err := db.Where("id=?",id).Delete(&user).Error; err!=nil {
		return err
	}
	return nil
}

// GetWholeNameList returns user name list
func (s * Service) GetWholeNameList() ([]string, error){
	db := database.GetDB()
	var names []string
	db.Table("users").Pluck("name", &names)


	return names,nil
}


