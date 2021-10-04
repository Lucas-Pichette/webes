package lib

import "fmt"

func criticalMsg(message string) string {
	return Style("bold", Style("red", "(!!) "+message))
}

// Display text in red and with an alert message-type specifier-text.
func errMsg(messagePieces ...string) string {
	var message string = "(!) "
	for _, messagePiece := range messagePieces {
		message += messagePiece + " "
	}
	// Style is visible in Styles.go, due to Styles.go being in the
	// same directory (lib) as this file (msgs.go), we're able to
	// access its exported files.
	return Style("r", message)
}

func warningMsg(message string) string {
	return Style("yellow", "(~) "+message)
}

func infoMsg(message string) string {
	return Style("gray", "(i) "+message)
}

func FmtPrint(message string, fmtTypes ...string) {
	// Print the now completely formatted message
	for _, fmtType := range fmtTypes {
		fmt.Println(Fmt(message, fmtType))
	}
}

func Fmt(message string, fmtType string) string {
	var formatSpecified bool = true

	// Run everything that changes the message itself
	switch fmtType {
	case "header", "head", "title":
		message = "== " + message + " =="
	default:
		formatSpecified = false
	}

	// Run the message-type specifier-text and Styleing.
	switch fmtType {
	case "critical", "c":
		message = criticalMsg(message)
	case "error", "err", "e":
		message = errMsg(message)
	case "warning", "w":
		message = warningMsg(message)
	case "info", "i":
		message = infoMsg(message)
	default:
		formatSpecified = false
	}

	// If there wasn't ever a fmt case that was met,
	// then inform the user that they are using the function
	// incorrectly.
	if !formatSpecified {
		errMsg("Incorrect fmtType specified")
	}
	return message
}
