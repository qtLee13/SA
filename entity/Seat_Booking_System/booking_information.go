package entity

import "time"

type Booking_Information struct {
	ID string `gorm:"primarykey"`
	DateAndTimeOfBooking time.Time
	PaymentStatus    string

	IDPassenger *uint
	IDPassenger Passenger `gorm:"foreignkey:IDPassenger;references:ID"`

	IDTrips *uint
	IDTrips Trips `gorm:"foreignkey:IDTrips;references:ID"`
}
