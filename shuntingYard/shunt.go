/*
 Package shuntingyard is used to implement the Shunting yard algorithim to vonvert a infix to Postix regular expression
*/

package shuntingyard

import (
	"bytes"
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
	var lastadded= ""
	// loop through the infix
	for curChar := 0; curChar < len(strArr); curChar++ {

		// retrive the last added value from the buffer
		strBuffer := []rune(buffer.String())

		// ensure buffer is not empty 
		if len(strBuffer) > 1 {
			 lastadded = string(strBuffer[len(strBuffer)-1])
		}

		// if it is the first character append to buffer.
		if charIsFirst(curChar) {
			buffer.WriteString(string(strArr[curChar]))
			continue
		}

		// if character is concatenator dont add it.
		if isConChar(string(strArr[curChar])) {
			continue
		}

		// if char is a special character  append to buffer.
		if isSpecials(string(strArr[curChar])) {

			buffer.WriteString(string(strArr[curChar]))
			continue
		}
		// if char is a opening bracket and not first char.
		if !charIsFirst(curChar) && isOpeningBracket(string(strArr[curChar])) {

			// if previouscharacter is concatenator dont add it.
			if isConChar(string(lastadded)) {
				buffer.WriteString(string(strArr[curChar]))
				continue
			} else {
				buffer.WriteString(".")
				buffer.WriteString(string(strArr[curChar]))
				continue
			}
		}

		// if char is not first and previous was   (
		if !charIsFirst(curChar) && isOpeningBracket(string(strArr[curChar-1])) {

			buffer.WriteString(string(strArr[curChar]))
			continue
		}

		//if character is closing bracket add to buffer
		if !charIsFirst(curChar) && isClosingBracket(string(strArr[curChar])) {

			buffer.WriteString(string(strArr[curChar]))
			continue
		}

		// If the previous character was not an OR | operater dont add concatante
		if !charIsFirst(curChar) && isOrOperator(string(strArr[curChar-1])) {
			buffer.WriteString(string(strArr[curChar]))
			continue
		}

		buffer.WriteString(".")
		buffer.WriteString(string(strArr[curChar]))
	}

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

func isOrOperator(char string) bool {
	if char == "|" {
		return true
	}
	return false
}
