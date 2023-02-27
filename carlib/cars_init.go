package carlib

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

var cars []Car

var errors int = 0
var initialized bool = false

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

// read cars data from the file into the 'cars' array
func initialize() {
	if initialized {
		return
	}

	cars = []Car{}

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

		cars = append(cars, car)
	}

	fmt.Printf("\n")
	fmt.Printf("Cars initialized! %d cars imported from file, %d errors occured.\n\n", count, errors)

	initialized = true

	return
}
