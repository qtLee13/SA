package entity

import "gorm.io/gorm"

type FeedBack struct {

	ID string `grom:"primarykey"`
	Name string
	Email    string
	Phonenumber	string
	FeedbackDetails string

	//IDbooking string
	//IDbooking Booking_Information `gorm:"foreignKey:IDbooking;references:ID"`

	//License_Plate string
	//License_Plate Vehicle `gorm:"foreignKey:License_Plate;references:License_Plate"`

	NameTypesOfFeedback TypeOfFeedback `gorm:"foreignKey:Name;references:Name"`
}
