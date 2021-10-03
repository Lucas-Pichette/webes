package main

import (
	"fmt"
	"os"

	"webes/lib"
)

type determineFunction func(string)

type Command struct {
	function    func()
	description string
}

var commands = make(map[string]Command)

func main() {
	commandHandler()
}

// Initializes a new webes project
// Callable via `webes init`
func webes_init() {
	lib.FmtPrint("initializing Project", "header", "info")
	const projectTree string = "" +
		"<pwd>\n" +
		"	┗━ dist/\n" +
		"		┣━ imgs/\n" +
		"		┣━ scripts/\n" +
		"		┣━ styles/\n" +
		"		┣━ pages/\n" +
		"		┗━ index.html\n" +
		"	┗━ dev/\n" +
		"		┣━ imgs/\n" +
		"		┣━ pages/\n" +
		"		┣━ components/\n" +
		"		┗━ app.go"

	makeProjectTree()

	// Now that we've made all of the directories, inform
	// the user of the changes.
	lib.FmtPrint("New Project with Directory Tree:", "info")
	fmt.Println(projectTree)
}

// Provides details about the various webes commands
// Callable via `webes help`
func webes_help() {
	lib.FmtPrint("Available Commands", "header", "info")
	for name, cmd := range commands {
		lib.FmtPrint(name+": "+cmd.description, "info")
	}
}

func makeProjectTree() {
	var paths = [7]string{"dist/imgs", "dist/scripts", "dist/styles", "dist/pages", "dev/imgs", "dev/pages", "dev/components"}
	for _, path := range paths {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}

	type fileT struct {
		path    string
		name    string
		content string
	}
	var files = [2]fileT{
		{
			path: "dist/",
			name: "index.html",
			content: "<!DOCTYPE HTML>\n<html lang='en-us'>" +
				"</html>\n<head>\n	<title></title>\n	<!--Metadata-->\n" +
				"	<meta charset='UTF-8'>\n	<meta name='viewport' content='width=device-width, initial-scale=1'>\n" +
				"	<meta name=robots content='index,follow'>\n	<meta name='description' content='!WEBSITE_DESCRIPTION'>\n" +
				"	<meta name='keywords' content='!RELEVANT_KEYWORDS'>\n	<meta name='author' content='!AUTHOR_NAME'>\n" +
				"	<meta rel='canonical' href='!YOUR_URL'>\n	<meta name='subject' content='!WEBSITE_SUBJECT'>\n" +
				"	<meta name='url' content='!YOUR_URL'>\n	<meta http-equiv='Expires' content='0'>\n" +
				"	<meta http-equiv='imagetoolbar' content='no'>\n\n	<!--Twitter Card Specs-->\n" +
				"	<meta name='twitter:card' content='summary'>\n	<meta name='twitter:site' content='@!YOUR_TWITTER_HANDLE'>\n" +
				"	<meta name='twitter:title' content='!YOUR_TITLE'>\n	<meta name='twitter:description' content='!YOUR_DESCRIPTION'>\n" +
				"	<meta name='twitter:image' content='!PATH_FOR_DISPLAY_IMAGE'>\n\n	<!--Open Graph Specs-->\n" +
				"	<meta property='og:title' content='!YOUR_TITLE'/>\n	<meta property='og:type' content='article'/>\n	<meta property='og:url' content='!YOUR_URL'/>\n" +
				"	<meta property='og:image' content='!PATH_FOR_DISPLAY_IMAGE'/>\n	<meta property='og:description' content='!YOUR_DESCRIPTION'/>\n" +
				"	<meta property='og:site_name' content='!YOUR_WEBSITE_NAME'/>\n\n	<!--Dependencies-->\n" +
				"	<link rel='stylesheet' href='!LINK_TO_YOUR_STYLESHEET'>\n</head>\n<body>\n	\n</body>\n</html>",
		},
		{
			path: "dev/",
			name: "app.go",
			content: "package dev\n\nimport \"fmt\"\n\nfunc app(){\n	fmt.Println(\"Hello, Webes!\")\n}\n",
		},
	}
	for _, file := range files {
		fileData := []byte(file.content)
		err := os.WriteFile(file.path+file.name, fileData, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func commandHandler() {
	args := os.Args[1:]
	runCommandInitializtion()

	if len(args) == 0 {
		lib.FmtPrint("Command not specified", "error")
		webes_help()
		return
	}

	var webes_command string = args[0]

	lib.FmtPrint("Running `"+webes_command+"`...", "info")

	if val, ok := commands[webes_command]; ok {
		// found command
		val.function()
	} else {
		// didn't find command
		lib.FmtPrint("Command not found", "error")
	}
}

// Function automatically ran during webes launch
func runCommandInitializtion() {
	commands["init"] = Command{
		function:    webes_init,
		description: "Initializes a new webes project",
	}
	commands["help"] = Command{
		function:    webes_help,
		description: "Provides details about the various webes commands",
	}
}
