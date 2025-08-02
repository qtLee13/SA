package entity

import "gorm.io/gorm"

type Seat struct {
	ID string `gorm:"primarykey"`
	SeatStatus string
	Email    string

	IDPassenger	*uint
	IDPassenger Passenger `gorm:"foreignkey:IDPassenger;references:ID"`

	IDTrips *uint
	IDTrips Trips `gorm:"foreignkey:IDTrips;references:ID"`
}
