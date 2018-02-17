package main

import (
    "fmt"
   
)

func stringPrinter(s string) {
    fmt.Printf(s)
}

func main() {
	var myString = "NFA String Matcher Project"
	stringPrinter(myString)
}