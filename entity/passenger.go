package entity


type Passenger struct {
	ID uint `gorm:"primarykey"`
	Name string
	Phonenumber string
	Email    string

}
