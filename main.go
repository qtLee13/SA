package main

import (
	"github.com/qtLee13/SA/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sa.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entity.FeedBack{}, &entity.TypeOfFeedback{}, &entity.Booking_Information{}, &entity.Passenger{}, &entity.Seat{}, &entity.Trips{})
}
