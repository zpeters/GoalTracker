package main

import (
	"fmt"
)

const goalsPath	= "c:\\users\\zach\\Projects\\ZachCore\\Organizer\\TODO.org"
const fitnessPath= "c:\\users\\zach\\Projects\\ZachCore\\Organizer\\Fitness.org"

const stretchesGoal	= true
const walkRunGoal	= 1.0
const squatsGoal	= 20
const pushupsGoal	= 10
const lungesGoal	= 20
const rowsGoal		= 10
const planksGoal	= 15
const jumpingJacksGoal	= 30
const weightGoal	= 190

type Goal struct {
	name			 string
	description		 string
	percent_complete	 float32
}

type Workout struct {
	date		 string
	stretches	 bool
	walkRun		 float64
	squats		 int64
	pushups		 int64
	lunges		 int64
	rows		 int64
	planks		 int64
	jumpingJacks	 int64
	weight		 float64 
}

func goalPrinter(g Goal) {
	fmt.Printf("\tGoal: %v\n", g)
}

func workoutPrinter(w Workout) {
	fmt.Printf("\tWorkout: %v\n", w)
}

func main() {
	fmt.Printf("Parse...\n")
	fmt.Printf("\tParse Long Term Goals\n")
	longTermGoal := Goal {}
	longTermGoal.name = "Long Term Goal"
	longTermGoal.description = "Long Term Goal Description"
	longTermGoal.percent_complete = 57.00

	fmt.Printf("\tParse Epic Goals\n")
	epicGoal := Goal {}
	epicGoal.name = "Epic Goal"
	epicGoal.description = "Epic Goal Description"
	epicGoal.percent_complete = 33.00

	fmt.Printf("\tParse Fitness\n")
	w1 := Workout{}
	w1.date = "2012-12-01"
	w1.stretches = true
	w1.walkRun = 1.3
	w1.squats = 10
	w1.pushups = 5
	w1.lunges = 4
	w1.planks = 20
	w1.jumpingJacks = 30
	w1.weight = 225.00
	fmt.Printf("\t\tYes/no: stretchs\n")
	fmt.Printf("\t\tFloat: walk, run\n")
	fmt.Printf("\t\tInt: squats, pushups, lunges, rows, plank, jumping jacks\n")
	fmt.Printf("\t\tFloat: weight\n")

	fmt.Printf("Execute template...\n")
	// just printing stuff for now
	goalPrinter(longTermGoal)
	goalPrinter(epicGoal)
	workoutPrinter(w1)
}