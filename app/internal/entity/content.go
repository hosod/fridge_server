package entity

import (
	"time"
)

// Content is contents of a fridge
// DBのテーブル定義用の構造体
type Content struct {
	ID int `json:"id"`
	ExpirationDate time.Time `json:"expiration_date" sql:"not null;type:date"`
	Quantity float32 `json:"quantity"`
	Fridge Fridge `json:"-" gorm:"foreignkey:FridgeID"`
	FridgeID int `json:"fridge_id" gorm:"not null"`
	FoodType FoodType `json:"-" gorm:"foreignkey:FoodTypeID"`
	FoodTypeID int `json:"food_type_id" gorm:"not null"`
}