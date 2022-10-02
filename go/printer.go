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

// cyan print cyan
func cyan(text string) {
	fmt.Println(Cyan + text + Reset)
}

// red print red
func red(text string) {
	fmt.Println(Red + text + Reset)
}

// green print green
func green(text string) {
	fmt.Println(Green + text + Reset)
}
