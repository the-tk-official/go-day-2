package main

import "log"

type User struct {
	// Struct of user
	FirstName string
}

func (user *User) printFirstName() {
	log.Println("My first name is", user.FirstName)
}

func main() {
	var firstUser User
	firstUser.FirstName = "Timmy"

	secondUser := User{
		FirstName: "Tony",
	}

	firstUser.printFirstName()
	secondUser.printFirstName()
}
