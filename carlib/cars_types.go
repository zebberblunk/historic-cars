package carlib

import "fmt"

type Car struct {
	name              string
	year              int
	country           string
	mileage           float64
	numberOfCylinders float64
	displacement      float64
	horsePowers       float64
	weight            float64
	acceleration      float64
	price             float64
}

func (car *Car) ToString(i int) (s string) {
	s = fmt.Sprintf("%d) %s (%s, %d) => %.0fhp, %.0fkg, %.0f cylinders | price: %.2f, displacement: %.1f, acceleration: %.1fs",
		i,
		car.name, car.country, car.year, car.horsePowers, car.weight,
		car.numberOfCylinders, car.price, car.displacement, car.acceleration)

	return
}
