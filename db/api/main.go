package main

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct {
	Make  string
	Model string
	Year  int
}

func (c Car) Drive() {
	fmt.Println("Driving", c.Make, c.Model, c.Year)
}

func main() {
	var car Vehicle = Car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,
	}
	car.Drive()
	fmt.Println("Hello, World!")
}
