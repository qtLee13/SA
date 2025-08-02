package entity

import "gorm.io/gorm"

type TypeOfFeedback struct {

	ID uint `gorm:"primaryKey"`
	Name string
	
}
