package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(gorilla)
}

func (d Dog) Says() string {
	return "Woof!"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "Grunt!"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	log.Println("The animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
