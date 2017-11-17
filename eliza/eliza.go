package eliza

import (
	"regexp"
)

//Creates new strucr
//code adapted from: https://gobyexample.com/structs
type Eliza_Response struct {
	re      *regexp.Regexp
	answers []string
}

//Function Respons uses paterns and array of answers
//Code adapted from: https://stackoverflow.com/questions/22282229/how-to-write-simple-regex-in-golang
//Code adapted from: https://shapeshed.com/golang-regexp/
func Response(pattern string, answers []string) Eliza_Response {
	response := Eliza_Response{}
	re := regexp.MustCompile(pattern)
	response.re = re
	response.answers = answers
	return response
} //End of Response function

func Ask(userInput string) string {

	return "your question:  " + userInput
} //End of ask function
