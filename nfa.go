/*
	Author		: Kevin Barry
	Start Date	: 22/09/2017
	Program		: build a non-deterministic finite automaton (NFA) from a regular expression and check it against any given string.
*/
package main

import (
	"fmt"
)

/*
	@param s a string to be printed
*/
func stringPrinter(s string) {
	fmt.Println(s)
}

/*

	@param s a string to be printed
	s is what is needed by the system

	@return UserInput a string entered by by user
*/
func userInput(s string) string {
	var userInput string
	stringPrinter(s)
	fmt.Scan(&userInput)
	return userInput
}

func main() {
	var myString = "NFA String Matcher Project"
	stringPrinter(myString)

	var regex = userInput("Enter Regex Expression: ")
	var testString = userInput("Enter String to Test: ")

	stringPrinter("Regex: " + regex + "\nString: " + testString)
}
