package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	myMap["me"] = User{
		FirstName: "Timur",
		LastName:  "Kurbanov",
	}

	log.Println(myMap["me"].FirstName)

	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println(mySlice)

	sort.Ints(mySlice)

	log.Println(mySlice)

	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	log.Println(numbers[5:])
}
