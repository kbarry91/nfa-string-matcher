/*
Package testCases is used to run tests on various parts of the program
*/
package testcases

import utils "../utils"
import shuntingYard "../shuntingYard"
import filereader "../filereader"


/*
TestMenu launches a menu to select a test*/
func TestMenu() {
	var testMenu = "============= Test  Menu  ============="
	returnToMain := false

	for !returnToMain {
		utils.StringPrinter("\n" + testMenu)
		progOption := utils.UserInput("1. Test String Concatenate \n2. Test Infix to Pofix \n3. Test String Input \n4. Test File Read \n5. Return To Main")

		switch progOption {
		case "1":
			TestConcat()
		case "2":
			TestInfixToPostfix()
		case "3":
			TestInput()
		case "4":
			TestFile()
		case "5":
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
	// DEBUG : ConcatenateInfix
	// TEST  : (ab)*(ab) should return (a.b)*.(a.b)
	// RESULT: Passedtest
	var testCat = utils.UserInput("DEBUG Concat: \n \tEnter String to Test: ")
	utils.StringPrinter("DEBUG Concat: \n \t" + testCat + " adapts to " + shuntingYard.ConcatenateInfix(testCat))
}

/*
TestInput testsUserInput Method
*/
func TestInput() {
	utils.StringPrinter("Test input Selected ....")

	// DEBUG : userInput
	// TEST  : test input by entering a string
	// RESULT: Passedtest
	var testString = utils.UserInput("DEBUG INPUT: \n \tEnter String to Test: ")
	utils.StringPrinter("DEBUG INPUT: \n \t " + testString)
}

/*
TestInfixToPostfix tests the shunting yard algorithim
*/
func TestInfixToPostfix() {
	utils.StringPrinter("test infix to pofix Selected ....")

	// DEBUG : InfixToPostfix
	// TEST  : test Converting an inﬁx regular expression (left) to postﬁx (right): a.(b.b)∗.a → abb.∗.a.
	// RESULT: Passedtest
	var shunter = shuntingYard.InfixToPostfix("a.(b.b)∗.a")
	utils.StringPrinter("DEBUG SHUNT:\n \t infix : a.(b.b)∗.a \n \t postfix :  " + shunter)
}

/*
TestFile tests the filereader
*/
func TestFile(){

	// Load a file
	fileWords := filereader.LoadDataFromFile()

	if len(fileWords) > 1{
		utils.StringPrinter("DEBUG File:\n \tFile read succesfull")

	}else{
		utils.StringPrinter("DEBUG File:\n \tFile read unsuccesfull")
	}

}