package eliza

import (
	"fmt"
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
func Response(patterns string, answers []string) Eliza_Response {
	response := Eliza_Response{}
	re := regexp.MustCompile(patterns)
	response.re = re
	response.answers = answers
	return response
} //End of Response function

//Function Ask compares user input if matches witch patterns
func Ask(userInput string) string {
	//This is one of the patern for testing purpose
	patters := "name is (.*)"
	// MustCompile, Compile to make a *regexp.Regexp struct
	re := regexp.MustCompile(patters)

	//If user input metched the pattern
	if re.MatchString(userInput) {

		fmt.Println("Match was found!")
		//Code adapted from: https://play.golang.org/p/YeSiBTfhFq
		match := re.FindStringSubmatch(userInput)

		//Match[1] mean the match was found
		found := match[1]
		fmt.Println(found)

		//Format string
		elizaOption := "Hello %s, How are you today?"
		//String formating
		//Code adapted from: https://gobyexample.com/string-formatting
		answer := fmt.Sprintf(elizaOption, found)
		fmt.Println(answer)

	} else {
		fmt.Println("No match found")
	}
	return ""

} //End of ask function
