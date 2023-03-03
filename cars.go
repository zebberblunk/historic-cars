package main

import (
	"fmt"

	"github.com/kjyrgen/historic-cars/carlib"
)

const DEFAULT_CHOICE string = "75"
const CHOICE_EXIT string = "0"

func main() {
	var choice string = DEFAULT_CHOICE

	for choice != CHOICE_EXIT {
		fmt.Println("\n= = = = = = = = = MAIN MENU = = = = = = = = = = = =")
		fmt.Println("")
		fmt.Println(" 0 => quit!\n")
		fmt.Println(" 1 => display full list of cars")
		fmt.Println(" 2 => search cars")
		fmt.Println(" 3 => display average car weights (1970 - 1982)")
		fmt.Println("")
		fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = =\n")
		fmt.Print("Your choice => ")
		fmt.Scanln(&choice)

		if choice == "1" {
			carlib.CarsDisplay()
		} else if choice == "2" {
			carlib.CarsSearch()
		} else if choice == "3" {
			carlib.DisplayBars()
		} else if choice == CHOICE_EXIT {
			fmt.Printf("\nProgram shutdown. Good night :)\n\n")
		} else {
			if choice != DEFAULT_CHOICE {
				fmt.Printf("\n\nThe option \"%s\" does not exist in the menu. Please try again.\n\n", choice)
			}
		}

		if choice != CHOICE_EXIT {
			choice = DEFAULT_CHOICE
		}
	}
}
