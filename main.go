package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	var users []User

	u1 := User{
		FirstName: "Timur",
	}

	u2 := User{
		FirstName: "Dilmurod",
	}

	users = append(users, u1)
	users = append(users, u2)

	for i, x := range users {
		log.Println(i, "-", x.FirstName)
	}

	for i := 0; i <= 2; i++ {
		log.Println(i)
	}

	mySlice := []string{
		"dog", "cat", "horse", "fish", "banana",
	}

	for i, x := range mySlice {
		log.Println(i, "-", x)
	}

	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	for x, i := range myMap {
		log.Println(i, "-", x)
	}
}
