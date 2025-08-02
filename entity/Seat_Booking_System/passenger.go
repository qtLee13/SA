package entity

import "gorm.io/gorm"

type Passenger struct {
	ID uint `gorm:"primarykey"`
	Name string
	Phonenumber string
	Email    string

}
