# Regular Expression To Non-deterministic Finite Automaton

A program written in Golang that can build a non-deterministic finite automaton (NFA) from a regular expression and check it against any given string.

### Synopsis

At the top of the file there should be a short introduction and/ or overview that explains **what** the project is. This description should match descriptions added for package managers (Gemspec, package.json, etc.)

### Motivation

This program was developed as a project for the Module 'Graph Theory'.

## Getting Started & Installing Required Softwares

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

In order to debug this Application you must have **Visual Studio Code** installed.
[How to install Visual Studio Code](https://code.visualstudio.com/download )

To run a Go program you must install the **Go Programming language** [How to install go](https://golang.org/dl/)

### How to Run and Debug the Project using Visual Studio Code

#####  Using the zip file
- Go to (https://github.com/kbarry91/nfa-string-matcher.git)
- Select to *Download Zip*
- Unzip the full folder
- Open visual studio code
- Navigate to 
> File > Open Folder..
- Select the unzipped solution

##### Alternativily using GIT
- Ensure Git is installed https://git-scm.com/downloads
- Create a new folder to hold the project
- Set up a git repository
- Navigate to directory of project in CMD
>Git init
>Git remote add origin https://github.com/kbarry91/nfa-string-matcher.git
>Git pull origin master

#### How to run the program
This program uses the Go programming language.
If you do not currently have Go installed click on the following link to download [INSTALL GO](https://golang.org/dl/)

To clone the repository to your local machine, in command prompt enter 
```
git clone https://github.com/kbarry91//nfa-string-matcher.git
```
There is two ways to run this program
1. **Build and Run** Navigate to the nfa-string-matcher folder and enter the following to compile the code 
```
go build nfa.go
```
This will create a .exe file in your current directory.To launch the program
```
nfa
```
2. **Run** to simply run the program in your command prompt enter the following 
```
go run nfa.go
```  

## Example Use

Show what the library does as concisely as possible, developers should be able to figure out **how** your project solves their problem by looking at the code example. Make sure the API you are showing off is obvious, and that your code is short and concise.

## Supported Regex Syntax

| Pattern                      | Description                                                 |
|:----------------------------:| ----------------------------------------------------------- |
| `a`, `b`, …                  | Simple match looking for the specified character            |
| `.`                          | Wildcard: match any character                               |
| `1\|0`                       | Alternation: match either `a` or `b`                        |
| `a*`                         | Repetition: match `a` zero or more times                    |
| `a+`                         | Repetition: match `a` one or more times                     |
| `a?`                         | Option: match `a` zero or one times                         |
| `(ab)`                       | Grouping                                                    |

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```



## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* Visual Studio Code
* Git

## Research 

Quick  paragraph abot how I researched the problem
see [References used for research](##acknowledgments-&-references)

## Authors

* **Kevin Barry** - *Initial work* - [kbarry91](https://github.com/kbarry91)

## License

This project is licensed under the 

(##acknowledgments-&-references)
* **Dr Ian McLoughlin** *lecturer of Graph Theory at G.M.I.T* [Ian McLoughlin](https://github.com/ianmcloughlin)

* Shunting yard algorithim : http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/#tocAnchor-1-7
* Regular Expression Matching : https://swtch.com/~rsc/regexp/regexp1.html
* Regex Cheat Sheat : http://www.rexegg.com/regex-quickstart.html
