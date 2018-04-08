# Regular Expression To Non-deterministic Finite Automaton

A program written in Golang that can build a non-deterministic finite automaton (NFA) from a regular expression and check it against any given string.

- [Regular Expression To Non-deterministic Finite Automaton](#regular-expression-to-non-deterministic-finite-automaton)
  * [Synopsis](#synopsis)
  * [Motivation](#motivation)
  * [Getting Started & Installing Required Softwares](#getting-started-and-installing-required-softwares)
    + [How to Run and Debug the Project using Visual Studio Code](#how-to-run-and-debug-the-project-using-visual-studio-code)
        * [Using the zip file](#using-the-zip-file)
        * [Alternativily using GIT](#alternativily-using-git)
    + [How to run the program](#how-to-run-the-program)
  * [Running the Tests](#running-the-tests)
    + [Test A Entered Expression Against an Entered String](#test-an-entered-expression-against-an-entered-string)
    + [Test A Entered Expression Against a File](#test-an-entered-expression-against-a-file)
  * [Operational Flow Of Program](#operational-flow-of-program)
  * [Supported Regex Syntax](#supported-regex-syntax)
  * [Built With](#built-with)
  * [Research & Design Process](#research-and-design-process)
  * [Limitiations Of The Program](#limitiations-of-the-program)
  * [Authors](#authors)
  * [Acknowledgments & References](#acknowledgments-and-references)


## Synopsis
The probem as stated in the the problem statement provided to us states _'You must write a program in the Go programming language [2] that can build a non-deterministic ﬁnite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. You must write the program from scratch and cannot use the regexp package from the Go standard library nor any other external library.'_

## Motivation

This program was developed as a project for the Module 'Graph Theory'. Graph Theory is a fundimental concept in Computer science and for that reason is added to the third year course work for Software Developement in GMIT Galway.

## Getting Started And Installing Required Softwares

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

### How to run the program
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

## Running the Tests
- The program has both a Launch Program and Test Mode. As running test mode is self explanitory I will describe how to run the main application. 
### Test An Entered Expression Against An Entered String
To allow future developement advancements the program has been developed with both a main program and test mode.
to run the program in main mode 

1. Launch the program as explained in How To Run a the Program.
2. A menu will appear with `1. Launch Program` or  `Launch test Mode` select `1. Launch Program`.
3. You are now in the String Matcher menu where ya can choose from `1. Match String` or `2. Find in File`. Select option 1.
4. Now you must enter a regular expression , enter `a?b?c`.
5. Next a string to check , enter `abc`
6. The program will now build a NFA from the expession and will display the following statistics
```
String Match Selected ....
This program has been configured to recognise an expression in infix or postfix notation
You may enter an expression in older or newer syntax (abc or a.b.c)
        Enter Regular Expression :
a?b?c?
        Enter String to Test:
abc
Entered regex    : a?b?c?
After Concatenate: a?.b?.c?
Postfix          : a?b?.c?.
        Match found for a?b?.c?. in abc !
```
### Test An Entered Expression Against a File

- It is important to note that all files must be stored in the '/data' folder and must have suffix `.txt`.
- This repository has been initilised with a extract from  _Deadlock_ by _Dorothy M. Richardson_ soucred from [Gutenburg](http://www.gutenberg.org/ebooks/56925m) 
- view the extract  [here](https://github.com/kbarry91/nfa-string-matcher/data/deadlock.txt])

1. Launch the program as explained above.
2. A menu will appear with `1. Launch Program` or  `Launch test Mode` select `1. Launch Program`.
3. You are now in the String Matcher menu where ya can choose from `1. Match String` or `2. Find in File`. Select option 2.
4. Enter a file Name without extension `deadlock`
5. Now you must enter a regular expression , enter `t.h.e` or `the`.
6. The program will now build a NFA from the expession and will display the following statistics
```
============= File Finder  =============
Enter File Name:  'deadlock' :
deadlock
File loading...
File loaded...
        Enter Regular Expression :
t.h.e
Entered regex    : t.h.e
After Concatenate: t.h.e
Postfix          : th.e.
Found 49 matches in file!
```
## Operational Flow Of Program
- Recieve an expression from user (infix or postix).
- Check expression syntax.
- If concatenator syntax not included adjust the statement.
- Check if expression is Infix or Postfix.
- If Infix , Convert to Postfix.
- Construct a NFA
- Check the NFA against a given String.
- Output statistics and result.

## Supported Regex Syntax
I have adjusted the program to adapt to the user input . That meaning if the expression entered is not in the correct syntax with or without the concatenator the program will catch this and make all the necessary changes. Aswell as this the program will convert an expression from infix to postfix notation 

| Input (infix notation)       | Concatenator added       | Postfix Notation       |                         
|:----------------------------:| ------------------------ | ---------------------- |
| `abc`                        | `a.b.c`                  | `ab.c`                 |
| `a.b.c`                      | `a.b.c`                  | `ab.c`                 |
| `abc(d|e)`                   | `a.b.c.(d|e)`            | `ab.c.de|.`            |
| `a?b?c?`                     | `a?.b?.c?`               | `a?b?.c?.`             |
| `00(0|1)*`                   |`0.0.(0|1)*`              | `00.01|*.`             |


| Pattern                      | Description                                                 |
|:----------------------------:| ----------------------------------------------------------- |
| `a`, `b`, …                  | Simple match looking for the specified character            |
| `.`                          | Wildcard: match any character                               |
| `1\|0`                       | Alternation: match either `a` or `b`                        |
| `a*`                         | Repetition: match `a` zero or more times                    |
| `a+`                         | Repetition: match `a` one or more times                     |
| `a?`                         | Option: match `a` zero or one times                         |
| `(ab)`                       | Grouping                                                    |


## Built With

* Visual Studio Code
* Git

## Research And Design Process

see [References used for research](##acknowledgments-and-references)

Initially on recieving this project I honestly felt it was out of the scope of my capabilities. For this reason I did not dive straight into the coding process.I  spent my first 2 weeks studying and researching about regular expressions and Finite Automota to ensure I had a clear understanding of what needed to be done. I found a brillant pdf online called **Introduction To The Theory Of Computation** by _Michael-Sipser*_ which can be downloaded [here](http://fuuu.be/polytech/INFOF408/Introduction-To-The-Theory-Of-Computation-Michael-Sipser.pdf)  which really goes into detail about the algorithims behind regular expresions.

The initial coding process began with following the videos posted on the Course module [see](#acknowledgments-and-references). This gave me a good benchmark to start the developement process. Once I understood this code I ran multiple tests to check for bugs and find areas to improve. At this point the program could parse an infix expression to postfix,construct a NFA and check a given string against it.

The first bug I found was that the program was not able to determine between an expression without the concatenator operator. This meaning although checking `a.b.c` against a string was succesfull ,`abc` would cause a fatal exception. After a more in depth reasearch of regular expressions, I could see a pattern and noted what caused a `.` to be added and on what conditions it would be left out. I designed and algorithim for the program to recognise if the concatenator was included or not and if not the expression would be adjusted.

I then realised that the program did not recognise enough patterns for it to be of use to the user. I added the `+ (match  one or more times)` and `? (match  zero or one times)` operators to the alorithim. This increased the amount of statements accepted ten fold.

The program was also adjusted so it can recognise a pattern that may contain any character , any letter [a-z,A-Z] and [0,9].

I decided to add file reading to improve the user expierence, I felt this was a very easy task as I just developed a seperate class to parse a file and take care of all the file reading. Once this was done the strings from the file ran through the program as normal.

## Limitiations Of The Program

In order to debug the program extensively and repetitively the program was developed with a Test Mode option when launched.This allowed me to test as much as possible in the quickest manner. I found that while entering and expression works perfectly, If an expression is copy and pasted into the program it causes inconsistencies in the result.

## Authors
* **Kevin Barry** - *Initial work* - [kbarry91](https://github.com/kbarry91)

## Acknowledgments And References
* **Dr Ian McLoughlin** *lecturer of Graph Theory at G.M.I.T* [Ian McLoughlin](https://github.com/ianmcloughlin)
* Shunting yard algorithim : http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/#tocAnchor-1-7
* Regular Expression Matching : https://swtch.com/~rsc/regexp/regexp1.html
* Regex Cheat Sheat : http://www.rexegg.com/regex-quickstart.html
* Introuction To The Theory Of computation : http://fuuu.be/polytech/INFOF408/Introduction-To-The-Theory-Of-Computation-Michael-Sipser.pdf
