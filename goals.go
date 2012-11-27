package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"html/template"
	"bytes"
)

const goalsPath	= "c:\\users\\zach\\Projects\\ZachCore\\Organizer\\TODO.org"
const fitnessPath = "c:\\users\\zach\\Projects\\ZachCore\\Organizer\\Fitness.org"
const templatePath = "Template"

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
	Name			 string
	Description		 string
	PercentComplete 	 float32
}

type Workout struct {
	Date		 string
	Stretches	 bool
	WalkRun		 float64
	Squats		 int64
	Pushups		 int64
	Lunges		 int64
	Rows		 int64
	Planks		 int64
	JumpingJacks	 int64
	Weight		 float64 
}

type Page struct {
	Goals []Goal
	Workouts []Workout
}

func goalParser() []Goal {
	// states
	state := "none"

	// goals array
	goals := []Goal{}

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
			g := Goal {
				Name: goalString,
				Description: state,
				PercentComplete: 0.0,
			}
			goals = append(goals, g)
		}
	}
	return goals
}

func main() {

	w1 := Workout{
		Date: "2012-12-01",
		Stretches:  true,
		WalkRun:  1.3,
		Squats:  10,
		Pushups:  5,
		Lunges:  4,
		Planks:  20,
		JumpingJacks:  30,
		Weight:  225.00,
	}

	
	goals := goalParser()
	workouts := []Workout{w1}

	t := template.New("Template")
	t, err := t.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	var out bytes.Buffer
	p := Page{
		Goals: goals,
		Workouts: workouts,
	}
	err = t.Execute(&out, p)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf(out.String())
}
