/*
	Author		: Kevin Barry
	Start Date	: 22/09/2017
	Program		: build a non-deterministic finite automaton (NFA) from a regular expression and check it against any given string.
*/
package main

import (
	"strconv"

	shuntingYard "./shuntingYard"
	testcases "./testcases"
	thompsons "./thompsons"
	utils "./utils"
)

func main() {
	var titleBar = "============= NFA String Main Menu ============="
	var subMenuBar = "============= String Matcher Menu  ============="
	//	var testMenu = "============= Test  Menu  ============="
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
			testcases.TestMenu()

			//TestMode()
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

//
// TestMode func is a main menu option used for developer debug
//
func TestMode() {
	/*

	   ------------TEST CODE------------

	*/

	// DEBUG :  Thompsons postfix to nfa and string matcher
	// TEST  :  Enter a string a.+b, test against abbb (should return true)
	// RESULT:  Test Passed
	utils.StringPrinter("DEBUG THOMP:")
	var testThomp = utils.UserInput("\t Enter Infix exp. :")
	var testAlpha = utils.UserInput("\t Enter Test str.  :")
	var isMatch = thompsons.StringMatcher(shuntingYard.InfixToPostfix(testThomp), testAlpha)
	utils.StringPrinter("\t Comparison       :" + strconv.FormatBool(isMatch))

}

//
// StringMatchMode allows user to enter a regular expression and test it aginst a string
//
func StringMatchMode() {
	utils.StringPrinter("This program has been configured to recognise an expression in infix or postfix notation")
	utils.StringPrinter("You may enter an expression in older or newer syntax (abc or a.b.c)")
	var userExp = utils.UserInput("\tEnter Regular Expression :")
	var userString = utils.UserInput("\tEnter String to Test: ")

	// Ensure expression is in correct syntax
	var conString = shuntingYard.ConcatenateInfix(userExp)

	// Convert expression from Infix to Postfix Notation
	var regexExpression = shuntingYard.InfixToPostfix(conString)

	// Construct a NFA from the expression and check to see if the expression is found in the String
	var isMatch = thompsons.StringMatcher(shuntingYard.InfixToPostfix(regexExpression), userString)

	// print some stats
	utils.StringPrinter("Entered regex    : " + userExp)
	utils.StringPrinter("After Concatenate: " + conString)
	utils.StringPrinter("Postfix          : " + regexExpression)
	switch isMatch {
	case true:
		utils.StringPrinter("\tMatch found for " + regexExpression + " in " + userString)
		break
	case false:
		utils.StringPrinter("\tNo Match found for " + regexExpression + " in " + userString)
		break
	default:
		utils.StringPrinter("\tAn error has occurred ,Please check regex expression and try again")
		break

	}

}
