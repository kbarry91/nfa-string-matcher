/*
Package filereader is used to parse data from a file
*/
package filereader

import ( utils "../utils"
 "io/ioutil"
 "strings"
)
/*
check is used to check for file read error
sourced from : https://gobyexample.com/reading-files
*/
func check(e error) {
	if e != nil {
		utils.StringPrinter("File Not found !")
	}
}

/*
LoadDataFromFile Loads data from a file and returns a slice of words
*/
func LoadDataFromFile() []string {

	filename := utils.UserInput("Enter File Name:  'deadlock' :")
	appendFileName :="data/"+filename+".txt"
	utils.StringPrinter("File loading...")
	
	// read from the file
	file, err := ioutil.ReadFile(appendFileName)
	check(err)

	// Parse the file to a string
	contents := string(file)

	// For each space spilt the string into words and add to slice
	words := strings.Split(contents, " ")
	return words
}
