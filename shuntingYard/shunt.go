/*
 Package shuntingyard is used to implement the Shunting yard algorithim to vonvert a infix to Postix regular expression
*/

package shuntingyard

import (
	"bytes"
	"fmt"

	utils "../utils"
)

/*
	InfixToPostfix implements the shunting yard algorithim and converts an expression from infix to postfix notion

	@param infix an expression in infux notation
	@returns a string postfix expresion
	adapted from http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/#tocAnchor-1-7
*/
func InfixToPostfix(infix string) string {

	prec := map[rune]int{'*': 10, '+': 9, '?': 8, '.': 6, '|': 5} //order of precedence of characters
	// * 0 or more
	// + 1 or more
	// ? Makes quantifiers "lazy"  as states at http://www.rexegg.com/regex-quickstart.html
	// . concatenate
	// | or

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

// ConcatenateInfix function  takes a infix expression and adds the concatenation "." where needed to make expression valid
// https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
func ConcatenateInfix(infix string) string {

	var buffer bytes.Buffer
	strArr := []rune(infix)

	// loop through the infix
	for curChar := 0; curChar < len(strArr); curChar++ {

		// if it is the first character append to buffer
		if charIsFirst(curChar) {
			buffer.WriteString(string(strArr[curChar]))
			continue
		}
		// if previous character is already concatenate operator
		if isConChar(string(strArr[curChar-1])) {
			buffer.WriteString(string(strArr[curChar]))
			continue
		}

		// if char is a special character append to buffer
		if isSpecials(string(strArr[curChar])) {
			buffer.WriteString(string(strArr[curChar]))
			continue
		}
		// if char is a opening bracket and not first char
		if !charIsFirst(curChar) && isOpeningBracket(string(strArr[curChar])) {
			buffer.WriteString(".")
			buffer.WriteString(string(strArr[curChar]))
			continue
		}
		// if char is not first and previous was not  (
		if !charIsFirst(curChar) && isOpeningBracket(string(strArr[curChar-1])) {

			buffer.WriteString(string(strArr[curChar]))
			continue
		}
		/*
			// if char is not first or last character and next is   )
			if !(curChar == len(infix)-1) {
				if !charIsFirst(curChar) && isClosingBracket(string(strArr[curChar])) {
					//	buffer.WriteString(".")
					buffer.WriteString(string(strArr[curChar]))
					continue
				}
			}

		*/
		//if character is closing bracket add to buffer
		if isClosingBracket(string(strArr[curChar])) {

			buffer.WriteString(string(strArr[curChar]))
			continue
		}
		buffer.WriteString(".")
		buffer.WriteString(string(strArr[curChar]))

	}

	fmt.Println(buffer.String())
	return buffer.String()
}

func charIsFirst(index int) bool {
	if index == 0 {
		return true
	}
	return false
}
func isOpeningBracket(char string) bool {
	if char == "(" {
		return true
	}
	return false
}
func isClosingBracket(char string) bool {
	if char == ")" {
		return true
	}
	return false
}

func isSpecials(char string) bool {

	specials := []string{"*", "+", "?", ".", "|"}

	for spec := range specials {
		if char == specials[spec] {
			utils.StringPrinter(char)
			return true

		}
	}
	return false
}
func isConChar(char string) bool {
	if char == "." {
		return true
	}
	return false
}
