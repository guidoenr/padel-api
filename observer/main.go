package main

import (
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"time"
)

const (
	BLINDEX = "https://darturnos.com/turnos/turnero/4188"
	CERRADA = "https://darturnos.com/turnos/turnero/4189"
)

type Observer struct {
	freq            time.Duration // frequency in minutes to observe
	startNotifyDate time.Time     // start time to send messages to the subs
	endNotifyDate   time.Time     // end time to send messages to the subs
	valuableHours   []string      // the hours to check for
}

type Turno struct {
	weekday string
	date    time.Time
	hour    string
}

type Config struct {
	Observer struct {
		Freq            int `json:"freq"`
		StartNotifyDate int `json:"startNotifyDate"`
		EndNotifyDate   int `json:"endNotifyDate"`
	} `json:"observer"`
	ValuableHours []string `json:"valuableHours"`
}

func (o *Observer) Initialize() {
	var config Config
	jsonBytes, _ := os.ReadFile("config.json")
	json.Unmarshal(jsonBytes, &config)

	o.freq = time.Duration(config.Observer.Freq) // TODO
	o.endNotifyDate = time.Now()                 // TODO
	o.startNotifyDate = time.Now()               // TODO , read right from the config
	o.valuableHours = config.ValuableHours

	fmt.Printf("initialized with freq: %d, valuableHours: %s, start: %v, end: %v", o.freq, o.valuableHours, o.startNotifyDate, o.endNotifyDate)
}

// GetTurnos obtain the turnos from the turnero, parse the response, and returns
// the most valuable turnos from the current datetime (time.Now)
func (o *Observer) GetTurnos(field string) {
	// getting the available turnos from the turneo
	response, err := soup.Get(field)
	if err != nil {
		fmt.Printf("error making request: %v", err)
	}

	// parse repsonse into HTMLParse
	doc := soup.HTMLParse(response)
	buttons := doc.FindAll("button")

	// getting today dates
	todayDate := time.Now()

	// parsing the data gotten from the website to the turno class
	var turnos []Turno
	for _, b := range buttons {
		weekday, date, hour := parseButton(b)
		if todayDate.Weekday().String() == weekday && isValuable(hour) {
			turnos = append(turnos, Turno{
				weekday: weekday,
				date:    date,
				hour:    hour,
			})
		}
	}

	for _, t := range turnos {
		fmt.Println(t)
	}
}

func isValuable(hour string) bool {
	return true
}

func main() {
	var obs Observer
	obs.Initialize()

	obs.GetTurnos(CERRADA)

}
