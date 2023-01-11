package main

import (
	"time"
)

const (
	BLINDEX = "https://darturnos.com/turnos/turnero/4188"
	CERRADA = "https://darturnos.com/turnos/turnero/4189"
)

func startObserving() {
	var obs Observer
	// start the observer
	if len(obs.valuableHours) == 0 {
		obs.Initialize()
	}

	var err error

	var turnosCerrada, turnosBlindex []Turno

	// cerrada turnos
	turnosCerrada, err = obs.GetTurnos(CERRADA, time.Now())
	if err != nil {
		time.Sleep(obs.freq)
		startObserving()
	}

	// blindex turnos
	turnosBlindex, err = obs.GetTurnos(BLINDEX, time.Now())
	if err != nil {
		time.Sleep(obs.freq)
		startObserving()
	}

	if len(turnosCerrada) == 0 && len(turnosBlindex) == 0 {
		time.Sleep(obs.freq)
		startObserving()
	} else {
		for _, t := range turnosBlindex {
			obs.turnosToDisplay = append(obs.turnosToDisplay, t)
		}
		for _, t := range turnosCerrada {
			obs.turnosToDisplay = append(obs.turnosToDisplay, t)
		}

		obs.NotifySubscribers()
		startObserving()
	}
}

func main() {

	startObserving()

}
