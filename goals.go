package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
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

func goalParser() []*Goal {
	// states
	state := "none"

	// goals array
	goals := []*Goal{}

	// regexes
	longTermRegex, err := regexp.Compile(`^\* Long Term`)
	epicRegex, err := regexp.Compile(`^\* Epic Goals`)
	goalRegex, err := regexp.Compile(`^\*\* `)

	// load the goals file
	content, err := ioutil.ReadFile(goalsPath)
	if err != nil { panic(err) }
	lines := strings.Split(string(content), "\n")
	if err != nil { panic(err) }

	for _ ,line := range lines {
		b := []byte(line)
		if longTermRegex.Match(b) {
			state = "longterm"
		} else if epicRegex.Match(b){
			state = "epic"
		} else if goalRegex.Match(b) {
			goalArray := strings.Split(line, "** ")
			goalString := goalArray[1]
			//fmt.Printf("Found an *%s* goal: %s\n", state, goalString)
			g := Goal {
				name: goalString,
				description: state,
				percent_complete: 0.0,
			}
			goals = append(goals, &g)
		}
	}
	return goals
}

func goalPrinter(g *Goal) {
	fmt.Printf("Name: %s\n", g.name)
	fmt.Printf("Description: %s\n", g.description)
	fmt.Printf("%g %%\n", g.percent_complete)
}

func workoutPrinter(w Workout) {
	fmt.Printf("Workout: %v\n", w)
}

func main() {

	// w1 := Workout{
	// 	date: "2012-12-01",
	// 	stretches:  true,
	// 	walkRun:  1.3,
	// 	squats:  10,
	// 	pushups:  5,
	// 	lunges:  4,
	// 	planks:  20,
	// 	jumpingJacks:  30,
	// 	weight:  225.00,
	// }

	//goalPrinter(longTermGoal)
	//goalPrinter(epicGoal)
	//workoutPrinter(w1)
	
	goals := goalParser()
	for _, goal := range goals {
		goalPrinter(goal)
		fmt.Println("")
	}
}
