package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Timur",
		LastName:  "Kurbanov",
	}

	log.Println(user.LastName, user.FirstName, user.BirthDate)
}
