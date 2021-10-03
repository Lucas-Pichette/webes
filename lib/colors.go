package lib

import "fmt"

var Reset string = "\033[0m"
var Red string = "\033[31m"
var Green string = "\033[32m"
var Yellow string = "\033[33m"
var Blue string = "\033[34m"
var Purple string = "\033[35m"
var Cyan string = "\033[36m"
var Gray string = "\033[37m"
var White string = "\033[97m"

// Capitalized functions are exportable
func Color(color string, text string) string {
	switch color {
	case "red", "r":
		return red(text)
	case "green", "g":
		return green(text)
	case "yellow", "y":
		return yellow(text)
	case "blue", "b":
		return blue(text)
	case "purple", "p":
		return purple(text)
	case "cyan", "c":
		return cyan(text)
	case "gray", "grey":
		return gray(text)
	case "white", "w":
		return white(text)
	}
	fmt.Println("Color not found, choose one of:\n" +
		"red, green, yellow, blue, purple, cyan, " +
		"gray, or white.")
	return text
}

func red(text string) string {
	return Red + text + Reset
}

func green(text string) string {
	return Green + text + Reset
}

func yellow(text string) string {
	return Yellow + text + Reset
}

func blue(text string) string {
	return Blue + text + Reset
}

func purple(text string) string {
	return Purple + text + Reset
}

func cyan(text string) string {
	return Cyan + text + Reset
}

func gray(text string) string {
	return Gray + text + Reset
}

func white(text string) string {
	return Red + text + Reset
}
