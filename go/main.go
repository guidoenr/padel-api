package main

import (
	"flag"
	"fmt"
	"github.com/anaskhan96/soup"
	"strconv"
	"strings"
	"time"
)

const (
	BLINDEX = "https://darturnos.com/turnos/turnero/4188"
	CERRADA = "https://darturnos.com/turnos/turnero/4189"
)

type Day struct {
	name  string
	hours []string
	date  string
}

func (d Day) addTurno(turno string) {
	d.hours = append(d.hours, turno)
}

func (d Day) showTurnos() {
	for _, h := range d.hours {
		fmt.Printf(h + "\n")
	}
}

func transformDate(date string) time.Time {
	d := strings.Split(date, "/")
	year, _ := strconv.Atoi(d[2])
	month, _ := strconv.Atoi(d[1])
	day, _ := strconv.Atoi(d[0])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, timeI)
}

func parseButton(turno soup.Root) (string, string, string) {
	// 4 index is the entire title description
	title := turno.Pointer.Attr[4].Val

	splitted := strings.Split(title, " ")

	// getting the turno's date and the available hour
	date := transformDate(splitted[5])
	hour := splitted[8]

	return date.Weekday().String(), date.String(), hour
}

func getTurnos(field string) []soup.Root {
	var url string
	if field == "blindex" {
		url = BLINDEX
	} else {
		url = CERRADA
	}

	// make the response
	response, err := soup.Get(url)
	if err != nil {
		fmt.Printf("error making request: %v", err)
	}

	// parse repsonse into HTMLParse
	doc := soup.HTMLParse(response)
	buttons := doc.FindAll("button")
	return buttons
}

func saveTurnos(field string) {
	var days []Day

	// all the turnos
	turnosCancha := getTurnos(field)

	// first turno
	auxDay, date, _ := parseButton(turnosCancha[0])
	fmt.Printf("day: %s date: %s", auxDay, date)
	// first day
	newDay := Day{
		name:  auxDay,
		hours: nil,
		date:  date,
	}

	for _, turno := range getTurnos(field) {
		day, date, hour := parseButton(turno)
		if auxDay == day {
			newDay.addTurno(hour)
		} else {
			days = append(days, newDay)
			newDay := Day{
				name:  day,
				hours: nil,
				date:  date,
			}
			newDay.addTurno(hour)
			auxDay = day
		}
	}

	for _, d := range days {
		d.showTurnos()
	}

	buttons := getTurnos(field)
	for _, b := range buttons {
		parseButton(b)
	}

}

func main() {
	field := flag.String("field", "blindex", "field name")
	flag.Parse()
	fmt.Println(*field)

	saveTurnos(*field)
}
