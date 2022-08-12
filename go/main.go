package main

import (
	"flag"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/guidoenr/vtools"
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
}

func parseButton(turno soup.Root) (string, string, string) {
	// 4 index is the entire title description
	title := turno.Pointer.Attr[4].Val

	splitted := strings.Split(title, " ")

	date := strings.Replace(splitted[5], "/", "-", 3)
	hour := splitted[8]

	layout := "13-02-2006"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("error parsing date: %v", err)
	}
	vtools.Print("parsedDate", t)

	return date, hour, date
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
	//var days []string

	// all the turnos
	turnosCancha := getTurnos(field)

	// first turno
	parseButton(turnosCancha[0])

	//# first day
	//new_day = Day(aux_day, date)
	//
	//for turno in turnos_cancha:
	//	day, date, hour = parse_button(turno)
	//if aux_day == day:
	//	new_day.add_turno(hour)
	//else:
	//	days.append(new_day)
	//	new_day = Day(day, date)
	//	new_day.add_turno(hour)
	//	aux_day = day
	//
	//for d in days:
	//	d.show_turnos()

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
