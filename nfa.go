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
	shunt implements the shunting yard algorithim and converts an expression from infix to postfix notion
	@param infix an expression in infux notation
	@returns a string postfix expresion
	adapted from http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/#tocAnchor-1-7
*/
func shunt(infix string) string {
	prec := map[rune]int{'*': 5, '.': 4, '|': 3} //order of precedence of characters

	postfix, stack := []rune{}, []rune{} //rune is a character as displayed on the screen(must convert back to string)

	// ( are added to stack
	// ) add each character stack to the postfix while the opening tag not found,remove character from stack
	// prec charcter , if it is less presidence value to tge last value add it to the postfix and remove from stack
	// (other characters eg 1 0 a b c ) add to postfix
	for _, rangeS := range infix { // range converts string to array of runes
		switch {
		case rangeS == '(':
			stack = append(stack, rangeS)
		case rangeS == ')':
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1] // retieve everything bar last element
			}
			stack = stack[:len(stack)-1]
		case prec[rangeS] > 0: // add precedence characters to stack after ()
			for len(stack) > 0 && prec[rangeS] <= prec[stack[len(stack)-1]] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1] // retieve everything bar last element
			}
			stack = append(stack, rangeS)
		default:
			postfix = append(postfix, rangeS)
		}
	}
	// if any characters remain on the stack append them to postix and clear the stack
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1] // retieve everything bar last element
	}

	return string(postfix)
}

func main() {
	var myString = "NFA String Matcher Project"
	stringPrinter(myString)
	var regex = userInput("Enter Regex Expression: ")

	/*

	   ------------TEST CODE------------

	*/
	//var testString = userInput("Enter String to Test: ")

	//  test carried out using below values and ran correctly
	//	test Converting an inﬁx regular expression (left) to postﬁx (right): a.(b.b)∗.a → abb.∗.a.
	stringPrinter("Regex: " + regex + "\nShunt :" + shunt(regex))
	//stringPrinter("Regex: " + regex + "\nString: " + testString)
}
