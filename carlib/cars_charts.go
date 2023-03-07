package carlib

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/pkg/browser"
)

// store all the car info into the map
// (year is the key, all the field values of that year are in array)
func getValuesMapForCountry(country string, field string) (carMap map[int][]float64) {
	carMap = make(map[int][]float64)
	for _, car := range cars {
		// country check
		if car.country != country {
			continue
		}

		// get the value
		f, ok := carMap[car.year]
		if !ok {
			f = []float64{}
		}

		value := car.GetValue(field)

		// if it is a 'displacement', convert the value from cubic inches to liters
		if field == DISPLACEMENT {
			value = value / 61
		}

		f = append(f, value)
		carMap[car.year] = f
	}

	return
}

// go through the values map and calculate/store yearly averages into separate map
func getMapOfAverages(valuesMap map[int][]float64) (mapOfAverages map[int]float64) {
	mapOfAverages = make(map[int]float64)
	for year, arrayOfValues := range valuesMap {
		var total float64 = 0
		for _, weight := range arrayOfValues {
			total += weight
		}
		average := total / float64(len(arrayOfValues))
		mapOfAverages[year] = average

		// fmt.Printf("Year %d, average %f\n", year, average)
	}

	return
}

// form a sorted array of yearly average values so we can use it on the chart
func getMapOfAveragesForChart(years []string, mapOfAverages map[int]float64) (chartValues []opts.BarData) {
	chartValues = make([]opts.BarData, len(years))

	for i, year := range years {
		y, _ := strconv.Atoi(year)
		yearlyAverage := mapOfAverages[y]
		chartValues[i] = opts.BarData{Value: yearlyAverage}

		// fmt.Printf("Bar values => year %s, average : %f\n", year, yearlyAverage)
	}

	return
}

func calculateYearlyAverages(country string, years []string, field string) (chartValues []opts.BarData) {
	valuesMap := getValuesMapForCountry(country, field)
	averagesMap := getMapOfAverages(valuesMap)
	chartValues = getMapOfAveragesForChart(years, averagesMap)

	return
}

func createBar(years []string, field string, titleExtension ...string) (bar *charts.Bar) {
	title := fmt.Sprintf("Average car %s by year", strings.ToLower(field))

	if len(titleExtension) > 0 {
		title = fmt.Sprintf("%s | %s", title, titleExtension[0])
	}

	// create a new bar instance
	bar = charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    title,
		Subtitle: "1970 - 1982 (blue USA; green Japan; yellow Europe)",
	}))

	// put data into the bar chart
	bar.SetXAxis(years).
		AddSeries("USA", calculateYearlyAverages("USA", years, field)).
		AddSeries("Japan", calculateYearlyAverages("Japan", years, field)).
		AddSeries("Europe", calculateYearlyAverages("Europe", years, field))

	return
}

func DisplayBars() {
	initialize()

	// create a map of all the "years" that the data covers (key: year)
	yearMap := make(map[string]int)
	for _, car := range cars {
		carYear := strconv.Itoa(car.year)

		// does the map already contain that year?
		_, ok := yearMap[carYear]
		// if not, add it!
		if !ok {
			yearMap[carYear] = car.year
		}
	}

	// convert it to string array of all the "years", and sort it ascendingly
	years := make([]string, len(yearMap))
	i := 0
	for year := range yearMap {
		years[i] = year
		i++
	}
	sort.Strings(years)

	// fmt.Printf("Years len %d; array %v\n", len(years), years)

	// create a new page and add the charts to it
	page := components.NewPage()
	page.AddCharts(
		createBar(years, WEIGHT),
		createBar(years, ACCELERATION),
		createBar(years, MILEAGE, "number of miles it can drive using one gallon of fuel"),
		createBar(years, HORSEPOWER),
		createBar(years, PRICE),
		createBar(years, CYLINDERS),
		createBar(years, DISPLACEMENT, "engine displacement in liters"),
	)

	// this is where the magic happens :)
	f, err := os.Create("temp/bars.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))

	// open the default browser with the generated chart
	browser.OpenFile("temp/bars.html")
}
