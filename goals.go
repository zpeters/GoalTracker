package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	"flag"
	"log"
)

const debug = false

var dataPath *string
var templatePath *string


type Goal struct {
	Name            string
	PercentComplete string
}

type Page struct {
	Goals     []Goal
	EpicGoals []Goal
	Studys    []Goal
	Todo      int
	Done      int
	PercentDone int
	Timestamp string
}

func dosOrUnixLineEndings(text string) string {
	dosNewLineRegex, _ := regexp.Compile(`\r\n`)
	if dosNewLineRegex.Match([]byte(text)) {
		return "dos" 
	} else {
		return "unix"
	}
	return "unknown"
}

func todoParser(files ...string) (todo, done int) {
	var todoCount int
	var doneCount int

	doneRegex,  _ := regexp.Compile(`^\** DONE`)
	checkedRegex, _ := regexp.Compile(`- [X]`)
	todoRegex, _ := regexp.Compile(`^\** TODO`)
	uncheckedRegex, _ := regexp.Compile(`\- \[ \]`)

	if debug { log.Printf("Processing todos/dones") }
	
	for _, file := range files {
		if debug { log.Printf("\tprocessing file %s", file) }
		content, _ := ioutil.ReadFile(*dataPath + file)
		lines := strings.Split(string(content), "\n")

		for _, line := range lines {
			b := []byte(line)
			if doneRegex.Match(b) || checkedRegex.Match(b) {
				if debug { log.Printf("\t\t%s", line) }
				if debug { log.Printf("\t\tfound done") }
				doneCount++
			} else if todoRegex.Match(b) || uncheckedRegex.Match(b) {
				if debug { log.Printf("\t\t%s", line) }		
				if debug { log.Printf("\t\tfound todo") }
				todoCount++
			}
		}
	}
	return todoCount, doneCount
}


func goalParser(files ...string) (normalGoals, epicGoals, studyGoals []Goal) {
	var state string = "none"

	// regexes
	normalRegex, _ := regexp.Compile(`^\* Long Term`)
	epicRegex, _ := regexp.Compile(`^\* Epic Goals`)
	studyRegex, _ := regexp.Compile(`^\* Study Goals`)
	goalRegex, _ := regexp.Compile(`^\*\*.+\[\d.%\]`)

	if debug { log.Printf("Processing Goals") }
	
	for _, file := range files {
		if debug { log.Printf("\tprocessing file %s", file) }
		var lines []string
		content, err := ioutil.ReadFile(*dataPath + file)
		if err != nil { panic(err) }

		// Never do this it's really dirty
		if dosOrUnixLineEndings(string(content)) == "dos" {
			lines = strings.Split(string(content), "\r\n")
		} else {
			lines = strings.Split(string(content), "\n")
		}

		for _, line := range lines {
			if debug { log.Printf("\t\t%s", line) }
			b := []byte(line)
			if normalRegex.Match(b) {
				if debug { log.Printf("\t\tState 'normal'") }
				state = "normal"
			} else if epicRegex.Match(b) {
				if debug { log.Printf("\t\tState 'epic'") }
				state = "epic"
			} else if studyRegex.Match(b) {
				if debug { log.Printf("\t\tState 'study'") }
				state = "study"
			}

			if goalRegex.Match(b) {
				if debug { log.Printf("\t\tFound goal") }
				lineArray := strings.Split(line, "** ")
				goalArray := strings.Split(lineArray[1], "[")
				goalString := goalArray[0]
				if debug { log.Printf("\t\tgoal string: %v", goalString) }

				percentArray := strings.Split(goalArray[1], "%]")
				percent := percentArray[0]
				if percent == "" {
					percent = "0"
				}

				g := Goal {
					Name: goalString,
					PercentComplete: percent,
				}

				if state == "normal" {
					normalGoals = append(normalGoals, g)
				} else if state == "epic" {
					epicGoals = append(epicGoals, g)
				} else if state == "study" {
					studyGoals = append(studyGoals, g)
				}
			}
		}
	}
	return normalGoals, epicGoals, studyGoals
	
}

func render() {
	var out bytes.Buffer

	normal, epic, study := goalParser("TODO.org", "Study.org")
	todo, done := todoParser("TODO.org", "TODO.org_archive")
	timestamp := time.Now().Format(time.ANSIC)

	t := template.New("Template")
	t, err := t.ParseFiles(*templatePath)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	p := Page{
		Goals:     normal,
		EpicGoals: epic,
		Studys:    study,
		Todo:      todo,
		Done:      done,
		PercentDone: (done * 100) / (todo + done),
		Timestamp: timestamp,
	}

	err = t.Execute(&out, p)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf(out.String())
}

func main() {
	dataPath = flag.String("dataPath", "Unknown", "Data path")
	templatePath = flag.String("templatePath", "Unknown", "Template path")
	flag.Parse()
	
	if (flag.NFlag() != 2) && (flag.NArg() != 2) {
		fmt.Printf("Usage:\n")
		fmt.Printf("-dataPath=/path/to/Organizer/\n")
		fmt.Printf("-templatePath=/path/to/Goals/Template\n")
		return
	} else {
		render()
	}

}
