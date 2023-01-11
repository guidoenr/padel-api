package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	BLINDEX = "https://darturnos.com/turnos/turnero/4188"
	CERRADA = "https://darturnos.com/turnos/turnero/4189"
)

type Day struct {
	name   string
	turnos []string
	date   time.Time
	field  string
}

// addTurno appends the turno in the day
func (d *Day) addTurno(turno string) {
	d.turnos = append(d.turnos, turno)
}

// showTurnos show all the turnos for the given day
func (d *Day) showTurnos() {
	green(fmt.Sprintf("[%s]-%d/%d", d.name, d.date.Day(), d.date.Month()))
	// barLength is the length of the bottom and top bars of the day
	barLength := 19
	fmt.Println(strings.Repeat("─", barLength))
	// what a nice algorithm to show the output in a pretty way , congrats monster
	// seeing if the index is multiple of 3, allow you to do an EndOfLine \n and print
	// the following set of 3 turnos
	for i, t := range d.turnos {
		if i%3 == 0 && i > 0 {
			fmt.Println("│")
		}
		fmt.Printf("│%s", t)
	}
	fmt.Print("│\n")
	fmt.Println(strings.Repeat("─", barLength))
}

// getTurnos do all the logic to get the available turnos for a given field
func getTurnos(field string) {
	// the days for the field
	var days []Day

	// getting all the turnos in the soup.Root type
	turnosCancha := requestTurnos(field)

	// first turno
	auxDay, auxDate, _ := parseButton(turnosCancha[0])
	// first day
	newDay := Day{
		field:  field,
		name:   strings.ToUpper(auxDay),
		turnos: nil,
		date:   auxDate,
	}

	for _, turno := range turnosCancha {
		day, date, hour := parseButton(turno)

		if auxDay == day {
			newDay.addTurno(hour)
		} else {
			days = append(days, newDay)
			newDay = Day{
				name: strings.ToUpper(day),
				date: date,
			}
			newDay.addTurno(hour)
			auxDay = day
		}
	}

	red(fmt.Sprintf("++++++[CANCHA]+++++"))
	red(fmt.Sprintf("+++++[%s]+++++", strings.ToUpper(field)))
	// for each day, show the turnos
	for _, d := range days {
		d.showTurnos()
	}

}

// getDate takes the date from the soup format and give you the time.Time datetime struct
func getDate(date string) time.Time {
	d := strings.Split(date, "/")
	year, _ := strconv.Atoi(d[2])
	month, _ := strconv.Atoi(d[1])
	day, _ := strconv.Atoi(d[0])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

// parseButton format the soup.Root struct and returns the weekday in string format, and the entire datetime as string
func parseButton(turno soup.Root) (string, time.Time, string) {
	// 4 index is the entire title description
	title := turno.Pointer.Attr[4].Val

	splitted := strings.Split(title, " ")

	// getting the turno's date and the available hour
	date := getDate(splitted[5])
	hour := splitted[8]

	return date.Weekday().String(), date, hour
}

// requestTurnos do the requests and get the responses
func requestTurnos(field string) []soup.Root {
	var url string
	if field == "blindex" {
		url = BLINDEX
	} else {
		url = CERRADA
	}

	// make the request and get the response
	response, err := soup.Get(url)
	if err != nil {
		fmt.Printf("error making request: %v", err)
	}

	// parse repsonse into HTMLParse
	doc := soup.HTMLParse(response)
	buttons := doc.FindAll("button")
	return buttons
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	fmt.Println("starting..")
	go routineGetTurnos(&wg, "blindex")
	go routineGetTurnos(&wg, "cerrada")

}

// routineGetTurnos is the routine for the get turnos func
func routineGetTurnos(wg *sync.WaitGroup, field string) {
	getTurnos(field)
	wg.Done()
}
