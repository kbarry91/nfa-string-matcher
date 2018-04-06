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
	var titleBar = "============= NFA String Main Menu ============="
	var subMenuBar = "============= String Matcher Menu  ============="
	//utils.StringPrinter(myString)

	// loop to keep main menu open
	keepRunning := true
	for keepRunning {
		utils.StringPrinter("\n" + titleBar)
		options := utils.UserInput("1. Launch Program\n2. Launch test Mode \n3. Exit")
		switch options {

		case "1":
			utils.StringPrinter("Launch Program Selected ....")
			returnToMain := false
			for !returnToMain {
				utils.StringPrinter("\n" + subMenuBar)
				progOption := utils.UserInput("1. Match String\n2. Match File \n3. Return To Main")
				switch progOption {
				case "1":
					utils.StringPrinter("String Match Selected ....")
					StringMatchMode()
				case "2":
					utils.StringPrinter("File Match Selected ....")
				case "3":
					utils.StringPrinter("Returning To main menu  ....")
					returnToMain = true
				default:
					utils.StringPrinter("Invalid Option")
					continue
				}
			}
			break
		case "2":
			utils.StringPrinter("Test Mode Selected ....")
			TestMode()
			break
		case "3":
			utils.StringPrinter("Exiting Program ....")
			keepRunning = false
			break
		default:
			utils.StringPrinter("Invalid Option")
			continue
		}

	}

}

/*
* TestMode func is a main menu option used for developer debug
 */
func TestMode() {
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

/*
* StringMatchMode allows user to enter a regular expression and test it aginst a string
 */
func StringMatchMode() {
	var userString = utils.UserInput("\tEnter String to Test: ")
	var userExp = utils.UserInput("\tEnter Regular Expression :")
	var isMatch = thompsons.StringMatcher(shuntingYard.InfixToPostfix(userExp), userString)
	switch isMatch {
	case true:
		utils.StringPrinter("\tMatch found for " + userExp + " in " + userString)
		break
	case false:
		utils.StringPrinter("\tNo Match found for " + userExp + " in " + userString)
		break
	default:
		utils.StringPrinter("\tAn error has occurred ,Please check regex expression and try again")
		break

	}

}
