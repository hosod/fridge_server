package entity

// Fridge is fridge info
type Fridge struct {
	ID int `json:"id"`
	Name string `json:"name" gorm:"not null"`
	User []*User `gorm:"many2many:user_fridge_relation;" json:"-"`
}