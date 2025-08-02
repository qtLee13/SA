package entity

import "time"

type Trips struct {
	ID uint `gorm:"primarykey"`
	DepartureTimestamp time.Time
	BasePrice int
	SeatCount int

	//RouteID *uint 
	//RouteID Routes `gorm:"foreignkey:RouteID;references:ID"`
	
	//VehicleID *uint
	//VehicleID Vehicle `gorm:"foreignkey:VehicleID;references:ID"`

	//OriginID *uint
	//OriginID Station `gorm:"foreignkey:OriginID;references:ID"`
	
	//DestinationID *uint
	//DestinationID Station `gorm:"foreignkey:DestinationID;references:ID"`

}
