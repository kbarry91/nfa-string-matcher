/*
	Author		: Kevin Barry
	Start Date	: 22/09/2017
	Program		: build a non-deterministic finite automaton (NFA) from a regular expression and check it against any given string.
*/
package main

import (
	shuntingyard "./shuntingyard"
	utils "./utils"
)

func main() {
	var myString = "NFA String Matcher Project"
	utils.StringPrinter(myString)
	var regex = utils.UserInput("Enter Regex Expression: ")
	//var testString = utils.UserInput("Enter string to test:")

	/*

	   ------------TEST CODE------------

	*/
	//var testString = userInput("Enter String to Test: ")

	// DEBUG : New shunting package
	utils.StringPrinter("DEBUG:" + "  infix: " + regex + " postfix :" + shuntingyard.InfixToPostfix(regex))
	//  test shunting yard algorithim
	//  test carried out using below values and ran correctly
	//	test Converting an inﬁx regular expression (left) to postﬁx (right): a.(b.b)∗.a → abb.∗.a.
	//utils.StringPrinter("  infix: " + regex + "\n  postfix :" + shunt(regex))
	//utils.StringPrinter("  infix: ." + regex + ". \n  postfix :." + shunt(regex) + ".")
	//utils.StringPrinter("testString ." + testString + ".")

	//isMatch := thompsons.StringMatcher(shunt(regex), testString)

	//utils.StringPrinter("  Does regex " + regex + ",po" + shunt(regex) + " match " + testString + " :" + strconv.FormatBool(isMatch))
}
