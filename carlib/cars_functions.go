package carlib

import (
	"fmt"
	"strings"
)

func CarsDisplay() {
	initialize()

	for i, car := range cars {
		fmt.Println(car.ToString(i + 1))
	}
}

func CarsSearch() {
	initialize()

	fmt.Print("\nPlease enter the name of the car OR country (usa, europe, japan) => ")
	var keyword string
	fmt.Scanln(&keyword)

	var found int = 0

	for _, car := range cars {
		var i int = strings.Index(strings.ToLower(car.name), strings.ToLower(keyword))
		var j int = strings.Index(strings.ToLower(car.country), strings.ToLower(keyword))

		if j >= 0 || i >= 0 {
			found++
			fmt.Println(car.ToString(found))
		}
	}
}
