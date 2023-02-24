package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"log"
	"strings"
)

const filename string = "data/cars-data.csv"

var errors int = 0

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

func readFloat(field string, index int, record []string) (f float64) {
	s := record[index]
	if s == "" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Car '%s'] Unable to read '%s': %v\n", record[0], field, err)

		errors++

		return -1
	}

	return
}

func main() {
	count := 0
	// read in the file into byte array
	buf, _ := ioutil.ReadFile(filename)

	// convert it into a string
	s := string(buf)

	// walk through the rows/records
	r := csv.NewReader(strings.NewReader(s))

	// .. and store the values into structs
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		count++

		// this is a header row
		if count == 1 {
			continue
		}

		year, _ := strconv.Atoi(record[1])
		mileage := readFloat("mileage", 3, record)
		numberOfCylinders := readFloat("numberOfCylinders", 4, record)
		displacement := readFloat("mileage", 5, record)
		horsePowers := readFloat("mileage", 6, record)
		weight := readFloat("mileage", 7, record)
		acceleration := readFloat("mileage", 8, record)
		price := readFloat("mileage", 9, record)

		var car Car = Car{
			name:              record[0],
			year:              year,
			country:           record[2],
			mileage:           mileage,
			numberOfCylinders: numberOfCylinders,
			displacement:      displacement,
			horsePowers:       horsePowers,
			weight:            weight,
			acceleration:      acceleration,
			price:             price,
		}

		fmt.Println(car.ToString(count - 1))
	}

	fmt.Printf("\n")
	fmt.Printf("Done! %d cars imported from file, %d errors occured.\n\n", count, errors)
}
