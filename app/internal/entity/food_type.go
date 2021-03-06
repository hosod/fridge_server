package entity

// FoodType is
type FoodType struct {
	ID      int       `json:"id"`
	Name    string    `json:"name" gorm:"unique;not null"`
	Image   string    `json:"image"`
	Genre   FoodGenre `json:"-" gorm:"foreignkey:GenreID"`
	GenreID int       `json:"-"`
	// Genre FoodGenre `json:"-" gorm:"foreignkey:GenreID"`
	// GenreID int `json:"-"`
}
