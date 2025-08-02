package entity

type Seat struct {
	ID string `gorm:"primarykey"`
	SeatStatus string
	Email    string

	IDPassenger	*uint
	Passenger Passenger `gorm:"foreignkey:IDPassenger;references:ID"`

	IDTrips *uint
	Trips Trips `gorm:"foreignkey:IDTrips;references:ID"`
}
