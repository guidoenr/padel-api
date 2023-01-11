package main

import (
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"time"
)

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

	fmt.Printf("Observer initialized with---->\nfrequency: %d\nvaluableHours: %s\nstartMessaging: %v\nendMessaging: %v\n", o.freq, o.valuableHours, o.startNotifyDate, o.endNotifyDate)
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
		turno := Turno{
			weekday:   weekday,
			date:      date,
			hour:      hour,
			field:     field,
			displayed: false,
		}
		if todayWeekDay == weekday && o.isValuable(turno) {
			turnos = append(turnos, turno)
		}
	}

	return turnos, nil
}

// isValuable will check if the given hour is valuable, looking into the valuable hours
// array that the observer has
func (o *Observer) isValuable(turno Turno) bool {
	for _, h := range o.valuableHours {
		if h == turno.hour {
			return true
		}
	}
	return false
}

// NotifySubscribers will notify all the subscribers
func (o *Observer) NotifySubscribers() {
	// TODO
	for _, t := range o.turnosToDisplay {
		fmt.Println(t.display())
	}
}
