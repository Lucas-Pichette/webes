# webes
_v0.0.3_  
  
**WEB**site **E**nvironment **S**ervice: Make websites in less time with Golang  
  
A website-development framework that **doesn't** hurt your SEO. All while 
allowing you to use your favorite features of a typical framework:  
* Components
  * Scoped styles and scripts
* Unused code and file cleanup upon project build
  
**TL;DR Webes cleans your environment upon `webes build` so that your website is lightning fast and SEO-efficient.**  
  
  
Development should occur in the dev/ directory, with the only exception being 
the index.html file in the dist/ directory.  

## Installation
Pre-Requisite: You must be in the directory you would like webes installed to.
Download via Git (assumes shell's PWD is in desired location for webes):  
```bash
git clone https://github.com/Lucas-Pichette/webes.git
```  

[comment]: <> (TODO: Add Installers for Each System)

## How to Use
To preview all of the possible webes-commands, enter:
`webes help`, or just `webes`. 
  
If you would prefer to use the interpreted version, as opposed to the 
executable, for any commands:
Run `go run main.go` in replacement of `webes`.  
  
To initialize/create a new project run:  
```bash
webes init
```  
  
The above command will create a directory tree that looks like:  
pwd  
&nbsp;&nbsp;&nbsp;&nbsp;┣━ dist/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┣━ imgs/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┣━ scripts/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┣━ styles/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┣━ pages/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┗━ index.html  
&nbsp;&nbsp;&nbsp;&nbsp;┗━ dev/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┣━ imgs/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┣━ pages/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┣━ components/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┗━ app.go  

## Versions
v0.0.4: Validation is Key!
* Updated main.go:
  * Added support for one new command `boilerplate`, added functionality to `validate`, and added an additional warning to `wipe`:
    * `boilerplate`: Once ran, will ask you for a file name, and then once said file name is provided, webes will generate a new HTML file with SEO-friendly boilerplate code in your PWD.
    * `validate` (WIP): Validate now finds unused code! Next up: Informating the user of file name, line number, and column number. ("Removing" unused code doesn't occur, it will leave out unused code when "building", which is a different command that will be made soon)
    * `wipe`: Added warning that informs user that `wipe` command must be ran in project-root directory. In other words, if you see div/ and dist/ when you ls, then you can run wipe successfully.

v0.0.3: Baby Steps
* Updated lib/colors.go:
  * Changed Color() function name to Style()
  * Added style options "bold" and "underline" to Style() function
* Updated lib/msgs.go:
  * Added criticalMsg(), which makes use of new "bold" style option
  * Refactored Fmt and FmtPrint
* Updated main.go:
  * Added support for two new commands:
    * `wipe`: Deletes the webes project that exists within the PWD.
    * `validate` (WIP): Checks dev/ for unused assets/code-segments.
  

v0.0.2: Webes Foundations  
* Created lib/ directory: Intended use is to hold Golang files, each file 
    representing some subject of utility functions.
* Created lib/colors.go: Intended use is to hold functions that relate to 
    Color. At the moment, all color-related functions in this file are for 
    terminal-based output.
* Created lib/msgs.go: Intended use is to hold functions that relate to 
    Messages. At the moment, all message-related functions in this file are for 
    terminal-based output.
* Created main.go: Intended use is to be the driver-file that manages 
    all-things related to the webes program. Currently manages two commands: 
    * `init`: Initializes a new webes project
    * `help`: Provides details about the various webes commands

