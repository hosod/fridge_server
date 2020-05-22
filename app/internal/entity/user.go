package entity

// import (
// 	"github.com/jinzhu/gorm"
// )

//User is user info
type User struct {
	// gorm.Model
	ID int `json:"id"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Fridge []*Fridge `gorm:"many2many:user_fridge_relation;" json:"-"`
}
