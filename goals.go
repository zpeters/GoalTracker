package main

import (
	"fmt"
)

const goalsPath = "c:\\users\\zach\\Projects\\ZachCore\\Organizer\\TODO.org"
const fitnessPath = "c:\\users\\zach\\Projects\\ZachCore\\Organizer\\Fitness.org"

type Goal struct {
	name string
	description string
	percent_complete float32
}

func main() {
	fmt.Printf("Parse...\n")
	fmt.Printf("Parse Long Term Goals: '%s'\n", goalsPath)
	fmt.Printf("Parse Epic Goals: '%s'\n", goalsPath)
	fmt.Printf("Parse Fitness: '%s'\n", fitnessPath)
	fmt.Printf("\tYes/no: stretchs\n")
	fmt.Printf("\tFloat: walk, run\n")
	fmt.Printf("\tInt: squats, pushups, lunges, rows, plank, jumping jacks\n")
	fmt.Printf("\tFloat: weight\n")
	fmt.Printf("Execute template\n")
}