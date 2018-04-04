/*
	Author		: Kevin Barry
	Start Date	: 22/09/2017
	Program		: build a non-deterministic finite automaton (NFA) from a regular expression and check it against any given string.
*/
package main

import (
	"strconv"

	shuntingYard "./shuntingYard"
	thompsons "./thompsons"
	utils "./utils"
)

func main() {
	var myString = "============= NFA String Matcher Project ============="
	utils.StringPrinter(myString)

	/*

	   ------------TEST CODE------------

	*/
	// DEBUG : userInput
	// TEST  : test input by entering a string
	// RESULT: Passedtest
	var testString = utils.UserInput("DEBUG INPUT: \n \tEnter String to Test: ")
	utils.StringPrinter("DEBUG INPUT: \n \t " + testString)

	// DEBUG : New shunting package
	// TEST  : test Converting an inﬁx regular expression (left) to postﬁx (right): a.(b.b)∗.a → abb.∗.a.
	// RESULT: Passedtest
	var shunter = shuntingYard.InfixToPostfix("a.(b.b)∗.a")
	utils.StringPrinter("DEBUG SHUNT:\n \t infix : a.(b.b)∗.a \n \t postfix :  " + shunter)

	// DEBUG :  Thompsons postfix to nfa and string matcher
	// TEST  :  Enter a string a.+b, test against abbb (should return true)
	// RESULT:  Test Passed
	utils.StringPrinter("DEBUG THOMP:")
	var testThomp = utils.UserInput("\t Enter Infix exp. :")
	var testAlpha = utils.UserInput("\t Enter Test str.  :")
	var isMatch = thompsons.StringMatcher(shuntingYard.InfixToPostfix(testThomp), testAlpha)
	utils.StringPrinter("\t Comparison       :" + strconv.FormatBool(isMatch))

}
