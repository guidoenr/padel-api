package main

import (
	"github.com/anaskhan96/soup"
	"strconv"
	"strings"
	"time"
)

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
