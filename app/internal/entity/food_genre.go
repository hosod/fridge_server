package entity

// FoodGenre is food-genre info
type FoodGenre struct {
	ID int `json:"id"`
	Name string `json:"name" gorm:"unique;not null"`
	Unit string `json:"unit"`
}