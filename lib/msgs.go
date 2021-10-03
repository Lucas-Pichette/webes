package lib

import "fmt"

// Display text in red and with an alert message-type specifier-text.
func errMsg(messagePieces ...string) string {
	var message string = "(!) "
	for _, messagePiece := range messagePieces {
		message += messagePiece + " "
	}
	// Color is visible in colors.go, due to colors.go being in the
	// same directory (lib) as this file (msgs.go), we're able to
	// access its exported files.
	return Color("r", message)
}

func infoMsg(message string) string {
	return Color("gray", "(i) "+message)
}

func FmtPrint(message string, fmtTypes ...string) {
	var formatSpecified bool = false

	// Run everything that changes the message itself
	for _, fmtType := range fmtTypes {
		switch fmtType {
		case "header", "head", "title":
			message = "== " + message + " =="
			formatSpecified = true
		}
	}

	// Run the message-type specifier-text and coloring.
	for _, fmtType := range fmtTypes {
		switch fmtType {
		case "error", "err", "e":
			message = errMsg(message)
			formatSpecified = true
		case "info", "i":
			message = infoMsg(message)
			formatSpecified = true
		}
	}

	// If there wasn't ever a fmt case that was met,
	// then inform the user that they are using the function
	// incorrectly.
	if !formatSpecified {
		errMsg("Incorrect fmtType specified")
		return
	}

	// Print the now completely formatted message
	fmt.Println(message)
}

func Fmt(message string, fmtTypes ...string) string {
	for _, fmtType := range fmtTypes {
		switch fmtType {
		case "error", "err", "e":
			return errMsg(message)
		case "info", "i":
			return infoMsg(message)
		case "header", "head", "title":
			return "== " + message + " =="
		}
	}
	errMsg("Incorrect fmtType specified")
	return message
}
