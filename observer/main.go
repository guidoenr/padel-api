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

type Turno struct {
	weekday   string    // the weekday in string
	date      time.Time // the date time of the turno
	hour      string    // the hour of the turno in format hh:mm (e.g 19:00)
	field     string    // the field, might be blindex/cerrada
	displayed bool      // displayed will set true if the turno was displayed
}

type Config struct {
	Observer struct {
		Freq            int `json:"freq"`
		StartNotifyDate int `json:"startNotifyDate"`
		EndNotifyDate   int `json:"endNotifyDate"`
	} `json:"observer"`
	ValuableHours []string `json:"valuableHours"`
}

type Observer struct {
	freq            time.Duration // frequency in minutes to observe
	turnosToDisplay []Turno       // the valuable turnos stored
	startNotifyDate time.Time     // start time to send messages to the subs
	endNotifyDate   time.Time     // end time to send messages to the subs
	valuableHours   []string      // the hours to check for
}

// Initialize will unmarshal the config read from the config.json file
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
// the most valuable turnos
func (o *Observer) GetTurnos(field string, dateToCheck time.Time) ([]Turno, error) {
	// getting the available turnos from the turneo
	response, err := soup.Get(field)
	if err != nil {
		return nil, fmt.Errorf("error getting response from the field %s: %v", field, err)
	}

	// parse repsonse into HTMLParse
	doc := soup.HTMLParse(response)
	buttons := doc.FindAll("button")

	// getting today dates
	todayWeekDay := dateToCheck.Weekday().String()

	// parsing the data gotten from the website to the turno class
	var turnos []Turno
	for _, b := range buttons {
		weekday, date, hour := parseButton(b)
		if todayWeekDay == weekday && o.isValuable(hour) {
			turnos = append(turnos, Turno{
				weekday: weekday,
				date:    date,
				hour:    hour,
				field:   field,
			})
		}
	}

	return turnos, nil
}

// isValuable will check if the given hour is valuable, looking into the valuable hours
// array that the observer has
func (o *Observer) isValuable(hour string) bool {
	for _, h := range o.valuableHours {
		if h == hour {
			return true
		}
	}
	return false
}

func main() {
	var obs Observer
	obs.Initialize()

	turnos, _ := obs.GetTurnos(BLINDEX, time.Now())
	fmt.Println(turnos)

}
