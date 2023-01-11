package main

import (
	"fmt"
	"time"
)

type Turno struct {
	weekday   string    // the weekday in string
	date      time.Time // the date time of the turno
	hour      string    // the hour of the turno in format hh:mm (e.g 19:00)
	field     string    // the field, might be blindex/cerrada
	displayed bool      // displayed will set true if the turno was displayed
}

func (t *Turno) display() string {
	msg := fmt.Sprintf("Turno available->\n DAY: %s\n FIELD: %s\n HOUR: %s\n", t.weekday, t.field, t.hour)
	return msg
}
