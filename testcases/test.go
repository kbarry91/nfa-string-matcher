/*
Package testCases is used to run tests on various parts of the program
*/
package testcases

import utils "../utils"
import shuntingYard "../shuntingYard"

/*
TestMenu launches a menu to select a test*/
func TestMenu() {
	var testMenu = "============= Test  Menu  ============="
	returnToMain := false
	for !returnToMain {
		utils.StringPrinter("\n" + testMenu)
		progOption := utils.UserInput("1. Test String Concatenate \n2. Test Infix to Pofix \n3. Return To Main")
		switch progOption {
		case "1":
			TestConcat()
		case "2":
			utils.StringPrinter("test infix to pofix Selected ....")
		case "3":
			utils.StringPrinter("Returning To main menu  ....")
			returnToMain = true
		default:
			utils.StringPrinter("Invalid Option")
			continue
		}
	}
}

/*
TestConcat tests the ConcatenateInfix method
*/
func TestConcat() {
	utils.StringPrinter("Test Concatenate Selected ....")
	// DEBUG : userInput
	// TEST  : test input by entering a string
	// RESULT: Passedtest
	var testCat = utils.UserInput("DEBUG Concat: \n \tEnter String to Test: ")
	utils.StringPrinter("DEBUG Concat: \n \t" + testCat + " adapts to " + shuntingYard.ConcatenateInfix(testCat))
}
