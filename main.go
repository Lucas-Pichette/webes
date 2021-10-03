package main

import (
	"fmt"
	"os"

	"webes/lib"
)

// The structure of a command that the user can execute.
type Command struct {
	function    func()
	description string
}

// The global commands-map that is filled in upon webes launch w/
// runCommandInitializtion().
var commands = make(map[string]Command)

func main() {
	commandHandler()
}

/* */
/* === WEBES COMMANDS === */
/* */
// Initializes a new webes project
// Callable via `webes init`
func webes_init() {
	lib.FmtPrint("initializing Project", "header", "info")
	const projectTree string = "" +
		"<pwd>\n" +
		"	┣━ dist/\n" +
		"		┣━ imgs/\n" +
		"		┣━ pages/\n" +
		"		┣━ scripts/\n" +
		"		┣━ styles/\n" +
		"		┗━ index.html\n" +
		"	┗━ dev/\n" +
		"		┣━ components/\n" +
		"		┣━ imgs/\n" +
		"		┣━ pages/\n" +
		"		┣━ scripts/\n" +
		"		┗━ styles/\n"

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

// Runs through dev/ directory and sub-directory validating HTML, CSS, and JS
// files to ensure that nothing exists that is not being used. Skips over
// comments.
// webes_validate automatically called when going to `webes build`.
func webes_validate() {

}

/* */
/* === Main-Level Functions === */
/* */
func commandHandler() {
	// Store the commandline arguments in a usable format
	args := os.Args[1:]

	// Initialize Commands
	runCommandInitializtion()

	// Determine if user forgot to specify a webes-command
	// (i.e. user entered `webes` as opposed to `webes init`)
	if len(args) == 0 {
		// If the user did forget to specify a webes-command,
		// then inform them that they did this, and provide
		// a list of options they could use, and exit the driver.
		lib.FmtPrint("Command not specified", "error")
		webes_help()
		return
	}

	var webes_command string = args[0]

	// It's useful to provide confirmation to the user, even if they don't
	// need it 99% of the time.
	lib.FmtPrint("Running `"+webes_command+"`...", "info")

	if val, ok := commands[webes_command]; ok {
		// found command
		val.function()
	} else {
		// didn't find command
		lib.FmtPrint("Command not found", "error")
	}
}

/* */
/* === Sub-Main-Level Functions === */
/* */
func makeProjectTree() {
	// Store all of the paths we want to create in the PWD that the command
	// `webes init` is called in.
	var paths = [9]string{
		"dist/imgs", "dist/scripts", "dist/styles", "dist/pages", "dev/imgs",
		"dev/pages", "dev/components", "dev/styles", "dev/scripts",
	}

	// For each specified path, attempt to create the full directory path,
	// and if there's an error, panic.
	for _, path := range paths {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}

	// The basic structure of a file-to-be-created.
	type fileT struct {
		path    string
		name    string
		content string
	}
	// Store all of the files that we want to create in the PWD
	var files = [4]fileT{
		{
			path: "dist/",
			name: "index.html",
			content: "<!DOCTYPE HTML>\n<html lang='en-us'>" +
				"</html>\n<head>\n	<title></title>\n	<!--Metadata-->\n	<" +
				"meta charset='UTF-8'>\n	<meta name='viewport' content='wi" +
				"dth=device-width, initial-scale=1'>\n	<meta name=robots con" +
				"tent='index,follow'>\n	<meta name='description' content='!WE" +
				"BSITE_DESCRIPTION'>\n	<meta name='keywords' content='!RELEV" +
				"ANT_KEYWORDS'>\n	<meta name='author' content='!AUTHOR_NAME" +
				"'>\n	<meta rel='canonical' href='!YOUR_URL'>\n	<meta nam" +
				"e='subject' content='!WEBSITE_SUBJECT'>\n	<meta name='url' " +
				"content='!YOUR_URL'>\n	<meta http-equiv='Expires' content='0" +
				"'>\n	<meta http-equiv='imagetoolbar' content='no'>\n\n	<" +
				"!--Twitter Card Specs-->\n	<meta name='twitter:card' content" +
				"='summary'>\n	<meta name='twitter:site' content='@!YOUR_TWI" +
				"TTER_HANDLE'>\n	<meta name='twitter:title' content='!YOUR" +
				"_TITLE'>\n	<meta name='twitter:description' content='!YOUR_D" +
				"ESCRIPTION'>\n	<meta name='twitter:image' content='!PATH_FOR" +
				"_DISPLAY_IMAGE'>\n\n	<!--Open Graph Specs-->\n	<meta pro" +
				"perty='og:title' content='!YOUR_TITLE'/>\n	<meta property='o" +
				"g:type' content='article'/>\n	<meta property='og:url' conte" +
				"nt='!YOUR_URL'/>\n	<meta property='og:image' content='!PATH_" +
				"FOR_DISPLAY_IMAGE'/>\n	<meta property='og:description' conte" +
				"nt='!YOUR_DESCRIPTION'/>\n	<meta property='og:site_name' con" +
				"tent='!YOUR_WEBSITE_NAME'/>\n\n	<!--Dependencies-->\n	<" +
				"link rel='stylesheet' href='styles/style.css'>\n</head>\n<bo" +
				"dy>\n	\n<!--Non-Critical Dependencies-->\n	<script langu" +
				"age=\"javascript\" type=\"text/javascript\" src=\"scripts/sc" +
				"ript.js\"></script>\n</body>\n</html>",
		},
		{
			path: "dev/styles/",
			name: "style.css",
			content: "html,body {\n	margin:0;\n	background-color:#333;\n	" +
				"color:white;\n}\n",
		},
		{
			path: "dev/components/",
			name: "_helloWorld.webes",
			content: "<template>\n	<div class='_helloWorld'>\n		<h1>" +
				"Hello, World!</h1>\n	</div>\n</template>\n\n\n<style>\n" +
				"	h1 {\n		font-size:250%;\n	}\n</style>\n\n\n" +
				"<script>\n	\n</script>\n",
		},
		{
			path:    "dev/scripts/",
			name:    "script.js",
			content: "console.log('Hello World!');\n",
		},
	}
	// For each file to be created, attempt to create said file with the
	// specified file data. If there's an error in this process, panic.
	for _, file := range files {
		fileData := []byte(file.content)
		err := os.WriteFile(file.path+file.name, fileData, 0644)
		if err != nil {
			panic(err)
		}
	}
}

// Function automatically ran during webes launch that ensures the
// commands map contains all of the commands that the user can execute.
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
