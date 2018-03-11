/*
Package utils is used to define basic util methds that will be used throughout the program
*/
package utils

import "fmt"

/*
StringPrinter prints a String s
*/
func StringPrinter(s string) {
	fmt.Println(s)
}

/*
UserInput takes an input from the user specified by s
	s is what is needed by the system

	@return UserInput a string entered by by user
*/
func UserInput(s string) string {
	var userInput string
	StringPrinter(s)
	fmt.Scan(&userInput)
	return userInput
}
