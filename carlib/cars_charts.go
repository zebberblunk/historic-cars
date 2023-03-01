package carlib

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/pkg/browser"
)

// generate random data for bar chart
func generateBarItems(amount int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < amount; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

// store all the car info into the map
// (year is the key, all the weight values of that year are in array)
func getWeightMapForCountry(country string) (carMap map[int][]float64) {
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

		f = append(f, car.weight)
		carMap[car.year] = f
	}

	return
}

// go through the car map and calculate/store yearly averages into separate map
func getAverageWeightMap(carMap map[int][]float64) (avgWeightMap map[int]float64) {
	avgWeightMap = make(map[int]float64)
	for year, weightArray := range carMap {
		var totalWeight float64 = 0
		for _, weight := range weightArray {
			totalWeight += weight
		}
		averageWeight := totalWeight / float64(len(weightArray))
		avgWeightMap[year] = averageWeight

		fmt.Printf("Year %d, average weight %f\n", year, averageWeight)
	}

	return
}

// form a sorted array of yearly average values so we can use it on the chart
func getYearlyWeightAveragesForChart(years []string, averageWeightMap map[int]float64) (chartValues []opts.BarData) {
	chartValues = make([]opts.BarData, len(years))

	for i, year := range years {
		y, _ := strconv.Atoi(year)
		yearlyAverageWeight := averageWeightMap[y]
		chartValues[i] = opts.BarData{Value: yearlyAverageWeight}

		fmt.Printf("bar values | year %s, average : %f\n", year, yearlyAverageWeight)
	}

	return
}

func calculateYearlyWeightAverages(country string, years []string) (chartValues []opts.BarData) {
	weightMap := getWeightMapForCountry(country)
	averageWeightMap := getAverageWeightMap(weightMap)
	chartValues = getYearlyWeightAveragesForChart(years, averageWeightMap)

	return
}

func DisplayCarWeightBar() {
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
			fmt.Printf("Year: %s\n", carYear)
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

	fmt.Printf("Years len %d; array %v\n", len(years), years)

	chartWeightAveragesUSA := calculateYearlyWeightAverages("USA", years)
	chartWeightAveragesJapan := calculateYearlyWeightAverages("Japan", years)
	chartWeightAveragesEurope := calculateYearlyWeightAverages("Europe", years)

	// create a new bar instance
	bar := charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Avereage car weights through the years (USA, Japan, Europe)",
		Subtitle: "1970 - 1982",
	}))

	// put data into the bar chart
	bar.SetXAxis(years).
		AddSeries("USA", chartWeightAveragesUSA).
		AddSeries("Japan", chartWeightAveragesJapan).
		AddSeries("Europe", chartWeightAveragesEurope)

	// this is where the magic happens :)
	f, _ := os.Create("data/bar.html")
	bar.Render(f)

	// open the default browser with the generated chart
	browser.OpenFile("data/bar.html")
}
