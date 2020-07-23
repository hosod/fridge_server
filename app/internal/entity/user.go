package entity

//User is user info
type User struct {
	// gorm.Model
	ID           int       `json:"id"`
	Name         string    `json:"name" gorm:"not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	FollowFridge []*Fridge `gorm:"many2many:user_follow_fridge;" json:"-"`
	MyFridge     Fridge    `json:"-" gorm:"foreignkey:MyFridgeID"`
	MyFridgeID   int       `json:"-"`
}
