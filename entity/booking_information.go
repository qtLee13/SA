package entity

import "time"

type Booking_Information struct {
	ID string `gorm:"primarykey"`
	DateAndTimeOfBooking time.Time
	PaymentStatus    string

	IDPassenger *uint
	Passenger Passenger `gorm:"foreignkey:IDPassenger;references:ID"`

	IDTrips *uint
	Trips   Trips `gorm:"foreignkey:IDTrips;references:ID"`
}
