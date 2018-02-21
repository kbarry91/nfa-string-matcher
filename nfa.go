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

/*
* shunt converts an expression from infix to postfix notion
* @param infix an expression in infux notation
* @returns a string postfix expresion
* adapted from http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/#tocAnchor-1-7
 */

func shunt(infix string) string {
	prec := map[rune]int{'*': 0, '.': 1, '|': 2} //order of precedence of characters
	postfix, stack := []rune{}, []rune{}         //rune is a character as displayed on the screen(must convert back to string)

	return string(postfix)
}
func main() {
	var myString = "NFA String Matcher Project"
	stringPrinter(myString)

	var regex = userInput("Enter Regex Expression: ")
	var testString = userInput("Enter String to Test: ")

	stringPrinter("Regex: " + regex + "\nString: " + testString + "\nShunt :" + shunt(regex))
	//stringPrinter("Regex: " + regex + "\nString: " + testString)
}
