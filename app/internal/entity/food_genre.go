package entity

// Fridge is fridge info
type FoodGenre struct {
	ID int `json:"id"`
	Name string `json:"name" gorm:"not null"`
	Unit string `json:"unit"`
}