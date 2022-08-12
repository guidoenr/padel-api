package main

import (
	"fmt"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// PrintTitle Print the input in a pretty way
func printTitle(text string) {
	fmt.Println(Cyan + "----[CANCHA " + text + "]----" + Reset)
}

func printHour(text string) {
	fmt.Printf("|   %s  |\n", text)
}

func printDay(text string) {
	fmt.Println(Green + "--[" + text + "]--" + Reset)
}
