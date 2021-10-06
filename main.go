package main

import (
	"fmt"       // Used for printing
	"io/ioutil" // Used for reading files and directories
	"os"        // Used for creating files and directories
	"strings"   // Used for string manipulation

	"webes/lib" // Used for various utility functions specific to webes
)

// The structure of a command that the user can execute.
type Command struct {
	function    func()
	description string
}

// 3 sections of a .webes file: <template>...,<style>..., and <script>...
// this struct will contain all the information for each of these sections
type parsedFileData struct {
	templateData parsedTemplateData
	styleData    parsedStyleData
	scriptData   parsedScriptData
}
type parsedTemplateData struct {
	classes []string
	ids     []string
	jsFuncs []string
}
type parsedStyleData struct {
	classes []string
	ids     []string
}
type parsedScriptData struct {
	jsFuncs []string
}

// The basic structure of a file-to-be-created.
type fileT struct {
	path    string
	name    string
	content string
}

// The global commands-map that is filled in upon webes launch w/
// runCommandInitializtion().
var commands = make(map[string]Command)
var pwd string

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("PWD: ", pwd)
	commandHandler()
}

/* */
/* === WEBES COMMANDS === */
/* */
// Creates a new boilerplate HTML file in PWD
func webes_boilerplate() {
	var boilerplate string = "<!DOCTYPE HTML>\n<html lang='en-us'>" +
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
		"dy>\n	\n\t<!--Non-Critical Dependencies-->\n	<script langu" +
		"age=\"javascript\" type=\"text/javascript\" src=\"scripts/sc" +
		"ript.js\"></script>\n</body>\n</html>"

	var fName string
	fmt.Scanln(&fName)

	fileData := []byte(boilerplate)
	if strings.Index(fName, ".html") == -1 {
		err := os.WriteFile(pwd+fName+".html", fileData, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		err := os.WriteFile(pwd+fName, fileData, 0644)
		if err != nil {
			panic(err)
		}
	}
}

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
		"			┗━ _helloWorld.webes\n" +
		"		┣━ imgs/\n" +
		"		┣━ pages/\n" +
		"		┣━ scripts/\n" +
		"			┗━ script.js\n" +
		"		┗━ styles/\n" +
		"			┗━ style.css\n"

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
	// 1) Scan through component files (*.webes)
	files, err := ioutil.ReadDir("dev/components")
	if err != nil {
		lib.FmtPrint(err.Error(), "error")
		panic(err.Error())
	}

	for _, f := range files {
		// The whole point of looking throught the files is to get data
		const maxFileSize int = 99999
		var fileData = make([]byte, maxFileSize)
		// and even further, to get the parsed data
		var pfd parsedFileData

		if f.IsDir() {
			fmt.Println("dev/components/" + f.Name())

			// Inform the user that at the moment they shouldn't have components
			// in sub-directories of dev/components/. That's a TODO for later.
			lib.FmtPrint("Directory within dev/components found with name: \""+
				f.Name()+"\", please ensure that you have all components in "+
				"dev/components, as opposed to nested within a sub-directory"+
				" within dev/components", "warning")
		} else {
			fmt.Println("dev/components/" + f.Name() + "/")

			var fileExtension = strings.SplitAfter(f.Name(), ".")[1]
			// Start the parsing if it's not a directory, and has *.webes format
			if fileExtension == "webes" {
				file, err := os.Open(pwd + "dev/components/" + f.Name())
				if err != nil {
					panic(err)
				} else {
					// if length is needed, replace _ w/ length
					_, err := file.Read(fileData)
					if err != nil {
						panic(err)
					} else {
						// read was successful
					}
				}
			}

			// Now that we have the file data, parse through it
			var fileStr string = string(fileData)
			scan(fileStr, "template", &pfd)
			scan(fileStr, "style", &pfd)
			scan(fileStr, "script", &pfd)
			fmt.Println("Completed all scans...")

			// 	1.d) Cross-compare what's used (collected from <template>...)
			//		with what exists (collected from <style>... and <script>...)

			// compare template classes with style classes
			for _, st_c := range pfd.styleData.classes {
				if !contains(pfd.templateData.classes, st_c) {
					// Add removal logic for that class in <style>
					// Add logic to determine file name, line number,
					// and column number as well
					lib.FmtPrint("Found unused class in style", "warning")
				}
			}
			// compare template ids with style ids
			for _, st_i := range pfd.styleData.ids {
				if contains(pfd.templateData.ids, st_i) {
					// Add removal logic for that id in <style>
					// Add logic to determine file name, line number,
					// and column number as well
					lib.FmtPrint("Found unused id in script", "warning")
				}
			}
			// compare template functions with script functions
			for _, sc_f := range pfd.scriptData.jsFuncs {
				if contains(pfd.templateData.jsFuncs, sc_f) {
					// Add removal logic for that function in <script>
					// Add logic to determine file name, line number,
					// and column number as well
					lib.FmtPrint("Found unused function in script", "warning")
				}
			}
			// compare style classes with template classes
			for _, t_c := range pfd.templateData.classes {
				if !contains(pfd.styleData.classes, t_c) {
					// Add removal logic for that class in <template>
					// Add logic to determine file name, line number,
					// and column number as well
					lib.FmtPrint("Found unused class in template", "warning")
				}
			}
			// compare style ids with template ids
			for _, t_i := range pfd.templateData.ids {
				if !contains(pfd.styleData.ids, t_i) {
					// Add removal logic for that id in <template>
					// Add logic to determine file name, line number,
					// and column number as well
					lib.FmtPrint("Found unused id in template", "warning")
				}
			}
			// compare script funtions with template functions
			for _, t_f := range pfd.templateData.jsFuncs {
				if !contains(pfd.scriptData.jsFuncs, t_f) {
					// Add removal logic for that function in <template>
					// Add logic to determine file name, line number,
					// and column number as well
					lib.FmtPrint("Found unused function in template", "warning")
				}
			}
		}
	}

}

func webes_wipe() {
	var userInput string

	lib.FmtPrint("FOR THIS COMMAND TO WORK YOU MUST BE IN THE ROOT OF YOUR "+
		"PROJECT (same directory as dev/ and dist/).", "warning")

	var confirmationMessage string = lib.Style("underline", "This is an "+
		"irreversible action.") +
		lib.Style("red", lib.Style("bold", " Are you sure that you want to "+
			"permanently delete this webes project (yes/no): "))
	confirmationMessage = lib.Fmt(confirmationMessage, "critical")

	// confirm with the user that they want to wipe the project
	fmt.Print(confirmationMessage)
	fmt.Scanln(&userInput)

	// if the user didn't enter an acceptable input, reprompt them until they do
	for strings.ToLower(userInput)[0] != 121 && // 121 == 'y'
		strings.ToLower(userInput)[0] != 110 { // 110 == 'n'
		fmt.Print(confirmationMessage)
		fmt.Scanln(&userInput)
	}

	if strings.ToLower(userInput)[0] == 121 {
		var pathToRemove = []string{"dev/", "dist/", "index.html"}

		for _, path := range pathToRemove {
			err := os.RemoveAll(pwd + path)
			if err != nil {
				panic(err)
			}
		}
	}
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
		// found commandv0.0.3: Baby Steps
		val.function()
	} else {
		// didn't find command
		lib.FmtPrint("Command not found", "error")
	}
}

/* */
/* === Sub-Main-Level Functions === */
/* */
func contains(s_arr []string, str string) bool {
	for _, e := range s_arr {
		if e == str {
			return true
		}
	}

	return false
}
func scan(fileStr string, whichScan string, pfd *parsedFileData) {
	// Remove filler characters created from []byte -> string casing
	actualFileSize := strings.Index(fileStr, "#end") +
		len("#end")
	fileStr = fileStr[:actualFileSize]
	// only retrieve info inbetween <template> and </template>
	startSection := strings.Index(fileStr, "<"+whichScan+">") +
		len("<"+whichScan+">")
	endSection := strings.Index(fileStr, "</"+whichScan+">")
	fileStr = fileStr[startSection:endSection]

	if whichScan == "template" {
		search(fileStr, "class=\"", "\"", &pfd.templateData.classes)
		search(fileStr, "class='", "'", &pfd.templateData.classes)

		search(fileStr, "id=\"", "\"", &pfd.templateData.ids)
		search(fileStr, "id='", "'", &pfd.templateData.ids)
	} else if whichScan == "style" {
		search(fileStr, ".", "{", &pfd.styleData.classes)
		search(fileStr, "#", "{", &pfd.styleData.ids)
	}
	for _, class := range pfd.templateData.classes {
		fmt.Println("templateData class: ", class)
	}
	for _, class := range pfd.styleData.classes {
		fmt.Println("styleData class: ", class)
	}
	for _, id := range pfd.templateData.ids {
		fmt.Println("templateData ids: ", id)
	}
	for _, id := range pfd.styleData.ids {
		fmt.Println("styleData ids: ", id)
	}

	if whichScan == "template" {
		var tokens = []string{"onclick", "onmouseover", "onmouseleave",
			"Offline", "Onabort", "onafterprint", "onbeforeonload",
			"onbeforeprint", "onblur", "oncanplay", "oncanplaythrough",
			"onchange", "oncontextmenu", "ondblclick", "ondrag",
			"ondragend", "ondragenter", "ondragleave", "ondragover",
			"ondragstart", "ondrop", "ondurationchange", "onemptied",
			"onended", "onerror", "onfocus", "onformchange", "onforminput",
			"onhaschange", "oninput", "oninvalid", "onkeydown",
			"onkeypress", "onkeyup", "onload", "onloadeddata",
			"onloadedmetadata", "onloadstart", "onmessage", "onmousedown",
			"onmousemove", "onmouseout", "onmouseover", "onmouseup",
			"onmousewheel", "onoffline", "ononline", "onpagehide",
			"onpageshow", "onpause", "onplay", "onplaying", "onpopstate",
			"onprogress", "onratechange", "onreadystatechange", "onredo",
			"onresize", "onscroll", "onseeked", "onseeking", "onselect",
			"onstalled", "onstorage", "onsubmit", "onsuspend",
			"ontimeupdate", "onundo", "onunload", "onvolumechange",
			"onwaiting"}
		for _, token := range tokens {
			search(fileStr, token+"=\"", "\"", &pfd.templateData.jsFuncs)
			search(fileStr, token+"='", "'", &pfd.templateData.jsFuncs)
		}
	} else if whichScan == "script" {
		search(fileStr, "function ", "{", &pfd.scriptData.jsFuncs)
		search(fileStr, "() => ", "{", &pfd.scriptData.jsFuncs)
	}
	for _, jsFunc := range pfd.templateData.jsFuncs {
		fmt.Println("jsFunc: ", jsFunc)
	}
	for _, jsFunc := range pfd.scriptData.jsFuncs {
		fmt.Println("jsFunc: ", jsFunc)
	}
}

func search(searchParameter string, startSubstr string, endSubstr string,
	arr *[]string) int {
	var resultsFound int = 0
	var className string
	var startIdx, endIdx int

	startIdx = strings.Index(searchParameter, startSubstr)
	// Ensure we're basing search on not-searched data
	if startIdx == -1 {
		return 0
	}
	endIdx = strings.Index(searchParameter[startIdx+len(startSubstr):], endSubstr)

	// Do search for classes via [class=", "]
	for startIdx != -1 && endIdx != -1 {
		// Store class name of previously parsed info
		className = searchParameter[startIdx+len(startSubstr) : startIdx+len(startSubstr)+endIdx]
		*arr = append(*arr, strings.TrimSpace(className))
		resultsFound++
		// update search parameter
		searchParameter = searchParameter[startIdx+len(startSubstr)+endIdx+len(endSubstr):]
		// restart search for new clas
		startIdx = strings.Index(searchParameter, startSubstr)
		if startIdx == len(searchParameter) ||
			startIdx == len(searchParameter)-1 {
			return resultsFound
		}
		// Ensure we're basing search on not-searched data
		endIdx = strings.Index(searchParameter[startIdx+len(startSubstr):], endSubstr)
	}
	return resultsFound
}

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
		err := os.MkdirAll(pwd+path, 0755)
		if err != nil {
			panic(err)
		}
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
				"dy>\n	\n\t<!--Non-Critical Dependencies-->\n	<script langu" +
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
			content: "<template>\n	<div id='title' class='_helloWorld'>\n" +
				"		<h1>Hello, World!</h1>\n	</div>\n</template>\n\n\n" +
				"<style>\n	h1 {\n		font-size:250%;\n	}\n	.unusedStyle " +
				"{\n		color:red;\n}\n</style>\n\n\n<script>\n	\n" +
				"</script>\n#end",
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
		err := os.WriteFile(pwd+file.path+file.name, fileData, 0644)
		if err != nil {
			panic(err)
		}
	}
}

// Function automatically ran during webes launch that ensures the
// commands map contains all of the commands that the user can execute.
func runCommandInitializtion() {
	commands["boilerplate"] = Command{
		function:    webes_boilerplate,
		description: "Creates a new boilerplate HTML file in PWD",
	}
	commands["init"] = Command{
		function:    webes_init,
		description: "Initializes a new webes project",
	}
	commands["help"] = Command{
		function:    webes_help,
		description: "Provides details about the various webes commands",
	}
	commands["validate"] = Command{
		function:    webes_validate,
		description: "(WIP) Checks dev/ for unused assets/code-segments.",
	}
	commands["wipe"] = Command{
		function:    webes_wipe,
		description: "Deletes the webes project that exists within the PWD.",
	}
}
