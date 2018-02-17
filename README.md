# Regular Expression To Non-deterministic Finite Automaton

A program written in Golang that can build a non-deterministic finite automaton (NFA) from a regular expression and check it against any given string.

### Synopsis

At the top of the file there should be a short introduction and/ or overview that explains **what** the project is. This description should match descriptions added for package managers (Gemspec, package.json, etc.)

### Motivation

This program was developed as a project for the Module 'Graph Theory'.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisities

What things you need to install the software and how to install them

```
Give examples
```

### Installing Required Softwares

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
or load doc/install.html in your web browser for installation
instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source or load doc/install-source.html
in your web browser for source installation instructions.

A step by step series of examples that tell you have to get a development env running

Stay what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Example Use

Show what the library does as concisely as possible, developers should be able to figure out **how** your project solves their problem by looking at the code example. Make sure the API you are showing off is obvious, and that your code is short and concise.

## Supported Regex Syntax

| Pattern                      | Description                                                 |
|:----------------------------:| ----------------------------------------------------------- |
| `1`, `0`, …                  | Simple match looking for the specified character            |
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

* Dropwizard - Bla bla bla
* Maven - Maybe
* Atom - ergaerga


## Research 

Quick  paragraph abot how I researched the problem

## Authors

* **Kevin Barry** - *Initial work* - [kbarry91](https://github.com/kbarry91)

## License

This project is licensed under the 

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc