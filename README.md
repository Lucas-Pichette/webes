# webes
_v0.0.3_  
  
**WEB**site **E**nvironment **S**ervice: Make websites in less time with Golang  
  
A website-development framework that **doesn't** hurt your SEO. All while 
allowing you to use your favorite features of a typical framework:  
* Components
  * Scoped styles and scripts
* Unused code and file cleanup upon project build
  
**TL;DR Webes cleans your environment upon `webes build` so that your website 
is lightning fast and SEO-efficient.**  
  
<details>
<summary>The best way to explain it is with an example:</summary>  
<br />
  
> _dev/_**Index.html**  
> ...  
> &lt;div class="my-unused-class"&gt;  
> &nbsp;&nbsp;&nbsp;&nbsp;&lt;_myComponent class="lg-2" contentVar=": )"&gt;
> Hello, Webes!&lt;/_myComponent&gt;  
> &lt;/div&gt;  
> ...  
> <br /> <br /> 
> _dev/components/_**_myComponent.webes**  
> &lt;template&gt;  
> &nbsp;&nbsp;&nbsp;&nbsp;&lt;div class="_myComponent"&gt;  
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&lt;h3&gt;{ innerText }
> &lt;/h3&gt;  
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&lt;button&gt;{ contentVar }
> &lt;/button&gt;  
> &nbsp;&nbsp;&nbsp;&nbsp;&lt;/div&gt;   
> &lt;/template&gt;  
>   
> &lt;style&gt;  
> &nbsp;&nbsp;&nbsp;&nbsp;p {  
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;color:blue;  
> &nbsp;&nbsp;&nbsp;&nbsp;}  
> &lt;/style&gt;  
>   
> &lt;script&gt;  
> &nbsp;&nbsp;&nbsp;&nbsp;window.addEventListener('DOMContentLoaded', function(){  
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;document.querySelectorAll
> ('button')&#91;0&#93;.style.color = 'red';  
> &nbsp;&nbsp;&nbsp;&nbsp;});  
> &lt;/script&gt;  
> <br /> <br /> 
> _dev/styles/_**Style.css**  
> ...  
> .lg-2{  
>     font-size:250%;  
> }  
> ...  
> <br /> <br /> 
> _dev/scripts/_**Script.js**  
> ...  
> function ususedFunc(){  
>     console.log("I have no purpose... ;(")  
> }  
> ...  
> <br /> <br /> 
> We can see in the above that in Index.html we have a component called 
> myComponent. In addition to this, we have some things that can be "cleaned". If 
> you run `webes validate`, webes will parse through your HTML, CSS, and JS files 
> and find everything not needed and remove it. This doesn't mean it will remove 
> it from your development files, but it will output messages in the terminal 
> warning you of where all of these unused code segments are.  
>   
> In otherwords, webes will go through your components (_myCompontent.webes) 
> first and will take the HTML and store it wherever &lt;_myComponent&gt;
> &lt;/_myComponent&gt; appears. Additionally, it will take the css and store it 
> in Style.css with a prefix of `._myComponent`. Notice in the _myComponent.webes 
> file that the `button` styling doesn't have that prefix already. That's because 
> webes will do that for you, in order to create a kind of "scoped" styling 
> environment for your component. Then it will take the code from the script and 
> add it to the Script.js file. Once it has done all of this, it will move onto 
> any other components. Once webes is done with all of the components, it moves 
> to checking if you have any unused images, or other assets. 
> </details>
  
Development should occur in the dev/ directory, with the only exception being 
the index.html file in the dist/ directory.  
  
<br /> <br /> 

## Installation
Pre-Requisite: You must be in the directory you would like webes installed to.
Download via Git (assumes shell's PWD is in desired location for webes):  
```bash
git clone https://github.com/Lucas-Pichette/webes.git
```  

[comment]: <> (TODO: Add Installers for Each System)
<br /> <br /> 

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
<br /> <br /> 

## Versions
  
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

