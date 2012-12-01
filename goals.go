package main

import (
	"bytes"
	"fmt"
	"log"
	"html/template"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const debug = false

const dataPath = "c:\\Users\\Zach\\Projects\\ZachCore\\Organizer\\"
//const dataPath = "/tmp/ZachCore/Organizer/"
//const dataPath = "/home/zach/Projects/ZachCore/Organizer/"
const templatePath = "Template"

const stretchesGoal = true
const walkRunGoal = 1.0
const squatsGoal = 20
const pushupsGoal = 10
const lungesGoal = 20
const rowsGoal = 10
const planksGoal = 15
const jumpingJacksGoal = 30
const weightGoal = 190

type Goal struct {
	Name            string
	Epic            bool
	PercentComplete string
}

type Study struct {
	Name            string
	PercentComplete string
}

type Workout struct {
	Date         string
	Stretches    bool
	WalkRun      float64
	Squats       int64
	Pushups      int64
	Lunges       int64
	Rows         int64
	Planks       int64
	JumpingJacks int64
	Weight       float64
}

type Page struct {
	Goals    []Goal
	EpicGoals    []Goal
	Workouts []Workout
	Studys    []Study
	Timestamp string
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
	content, err := ioutil.ReadFile(dataPath + "TODO.org")
	if debug { log.Printf("Parsing file %s", dataPath + "TODO.org") }
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if err != nil {
		panic(err)
	}

	if debug { log.Printf("Found %d lines", len(lines)) }

	for i, line := range lines {
		if debug { log.Printf("\tProcessing line: %s", line) }
		if debug { log.Printf("\tProcessing line: %d", i) }
		b := []byte(line)
		if longTermRegex.Match(b) {
			state = "longterm"
		} else if epicRegex.Match(b) {
			state = "epic"
		} else if goalRegex.Match(b) && state == "longterm" {
			goalArray := strings.Split(line, "** ")
			myArray := strings.Split(goalArray[1], "[")
			goalString := myArray[0]
			if debug { log.Printf("\t\tNormal goal: %s", goalString) }
			percentArr := strings.Split(myArray[1], "%]")
			percentString := strings.Split(percentArr[0],"[")
			percent := percentString[0]
			if percent == "" {
				percent = "0"
			}
			if debug { log.Printf("\t\t\tPercent Complete: %s", percent) }
			
			var epic bool
			if state == "epic" {
				epic = true 
			} else {
				epic = false
			}
			
			g := Goal{
				Name:            goalString,
				Epic:            epic,
				PercentComplete: percent,
			}
			goals = append(goals, g)
		} else {
			if debug { log.Printf("\t\t Ignored...") }
		}
	}
	return goals
}

func goalParserEpic() []Goal {
	// states
	state := "none"

	// goals array
	goals := []Goal{}

	// regexes
	longTermRegex, err := regexp.Compile(`^\* Long Term`)
	epicRegex, err := regexp.Compile(`^\* Epic Goals`)
	goalRegex, err := regexp.Compile(`^\*\* `)

	// load the goals file
	content, err := ioutil.ReadFile(dataPath + "TODO.org")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		b := []byte(line)
		if longTermRegex.Match(b) {
			state = "longterm"
		} else if epicRegex.Match(b) {
			state = "epic"
		} else if goalRegex.Match(b) && state == "epic" {
			goalArray := strings.Split(line, "** ")
			myArray := strings.Split(goalArray[1], "[")
			if debug { log.Printf("My Array: %v", myArray) }
			goalString := myArray[0]
			percentArr := strings.Split(myArray[1], "%]")
			percentString := strings.Split(percentArr[0],"[")
			percent := percentString[0]
			if percent == "" {
				percent = "0"
			}

			var epic bool
			if state == "epic" {
				epic = true 
			} else {
				epic = false
			}
			
			g := Goal{
				Name:            goalString,
				Epic:            epic,
				PercentComplete: percent,
			}
			goals = append(goals, g)
		}
	}
	return goals
}

func studyParser() []Study {
	// goals array
	goals := []Study{}

	// regexes
	goalRegex, err := regexp.Compile(`^\*\* `)

	// load the goals file
	content, err := ioutil.ReadFile(dataPath + "Study.org")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		b := []byte(line)
		if goalRegex.Match(b) {
			goalArray := strings.Split(line, "** ")
			myArray := strings.Split(goalArray[1], "[")
			goalString := myArray[0]
			percentArr := strings.Split(myArray[1], "%]")
			percentString := strings.Split(percentArr[0],"[")
			percent := percentString[0]
			if percent == "" {
				percent = "0"
			}

			g := Study{
				Name:            goalString,
				PercentComplete: percent,
			}
			goals = append(goals, g)
		}
	}
	return goals
}

func workoutParser() []Workout {
	// workouts array
	workouts := []Workout{}

	// regexes
	workoutRegex, err := regexp.Compile(`^\| <`)

	// load the workouts file
	content, err := ioutil.ReadFile(dataPath + "Fitness.org")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		b := []byte(line)
		if workoutRegex.Match(b) {
			workoutArray := strings.Split(line, "|")

			date := workoutArray[1]
			stretches, err := strconv.ParseBool(strings.TrimSpace(workoutArray[2]))
			if err != nil {
				panic(err)
			}
			walkRun, err := strconv.ParseFloat(strings.TrimSpace(workoutArray[3]), 64)
			if err != nil {
				panic(err)
			}
			squats, err := strconv.ParseInt(strings.TrimSpace(workoutArray[4]), 10, 64)
			if err != nil {
				panic(err)
			}
			pushups, err := strconv.ParseInt(strings.TrimSpace(workoutArray[5]), 10, 64)
			if err != nil {
				panic(err)
			}
			lunges, err := strconv.ParseInt(strings.TrimSpace(workoutArray[6]), 10, 64)
			if err != nil {
				panic(err)
			}
			rows, err := strconv.ParseInt(strings.TrimSpace(workoutArray[7]), 10, 64)
			if err != nil {
				panic(err)
			}
			planks, err := strconv.ParseInt(strings.TrimSpace(workoutArray[8]), 10, 64)
			if err != nil {
				panic(err)
			}
			jumpingJacks, err := strconv.ParseInt(strings.TrimSpace(workoutArray[9]), 10, 64)
			if err != nil {
				panic(err)
			}
			weight, err := strconv.ParseFloat(strings.TrimSpace(workoutArray[10]), 64)
			if err != nil {
				panic(err)
			}

			w := Workout{
				Date:         date,
				Stretches:    stretches,
				WalkRun:      walkRun,
				Squats:       squats,
				Pushups:      pushups,
				Lunges:       lunges,
				Rows:         rows,
				Planks:       planks,
				JumpingJacks: jumpingJacks,
				Weight:       weight,
			}

			workouts = append(workouts, w)
		}
	}
	return workouts
}

func printOut() {
	var out bytes.Buffer

	goals := goalParser()
	epicGoals := goalParserEpic()
	study := studyParser()
	workouts := workoutParser()
	timestamp := time.Now().Format(time.ANSIC)

	t := template.New("Template")
	t, err := t.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	p := Page{
		Goals:    goals,
		EpicGoals: epicGoals,
		Studys:    study,
		Workouts: workouts,
		Timestamp: timestamp,
	}

	err = t.Execute(&out, p)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf(out.String())
}

func main() {
	printOut()
}
