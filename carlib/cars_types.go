package carlib

import "fmt"

const WEIGHT = "weight"
const HORSEPOWER = "horsepower"
const ACCELERATION = "acceleration"
const DISPLACEMENT = "displacement"
const PRICE = "price"
const CYLINDERS = "cylinders"
const MILEAGE = "mileage"

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

func (car *Car) GetValue(fieldName string) (value float64) {
	if fieldName == WEIGHT {
		value = car.weight
	} else if fieldName == HORSEPOWER {
		value = car.horsePowers
	} else if fieldName == ACCELERATION {
		value = car.acceleration
	} else if fieldName == MILEAGE {
		value = car.mileage
	} else if fieldName == PRICE {
		value = car.price
	} else if fieldName == CYLINDERS {
		value = car.numberOfCylinders
	} else if fieldName == DISPLACEMENT {
		value = car.displacement
	} else {
		panic(fmt.Sprintf("Unknown field name %s", fieldName))
	}

	return
}

func (car *Car) ToString(i int) (s string) {
	s = fmt.Sprintf("%d) %s (%s, %d) => %.0fhp, %.0fkg, %.0f cylinders | price: %.2f, displacement: %.1f, acceleration: %.1fs",
		i,
		car.name, car.country, car.year, car.horsePowers, car.weight,
		car.numberOfCylinders, car.price, car.displacement, car.acceleration)

	return
}
