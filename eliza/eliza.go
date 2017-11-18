package eliza

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
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

	allPatterns := docPatterns()

	//Looping true all patterns and checks if a match is found
	for _, elizaPattern := range allPatterns {

		//Code adapted from: https://golang.org/pkg/regexp/#example_Regexp_MatchString
		if elizaPattern.re.MatchString(userInput) {
			//Code adapted from: https://golang.org/pkg/regexp/#example_Regexp_FindStringSubmatch
			match := elizaPattern.re.FindStringSubmatch(userInput)
			found := match[1]
			//Reflect function
			found = Reflect(found)
			//Takes random pattern
			formatAnswer := randomAnswer(elizaPattern.answers)
			//Format string
			if strings.Contains(formatAnswer, "%s") {
				formatAnswer = fmt.Sprintf(formatAnswer, found)
			}
			return formatAnswer

		}

	} //End of ask function
	/*//This is one of the patern for testing purpose
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
	*/
	return "Sorry but I don't understand"
}

//This function builds list of responses
func docPatterns() []Eliza_Response {

	allResponses := []Eliza_Response{}

	//Code adapted from: Book(An Introdution to Programming in Go p.136)
	file, err := os.Open("./doc/patterns.dat")
	// handle the error here return
	if err != nil {
		//Crash here
		panic(err)
	}
	defer file.Close()

	//Code adapted from :https://gobyexample.com/line-filters
	//Wrapps the unbuffered file with a buffered scanner and gives a convenient Scan method that advances the scanner to the next token; which is the next line in the default scanner.
	scanner := bufio.NewScanner(file)

	//Text returns the current token, here the next line, from the input.
	for scanner.Scan() {
		pattern := scanner.Text()
		scanner.Scan()
		elizaAnswer := scanner.Text()
		//; is used in patterns.dot to splis list of eliza answers
		listOfelizaAnswer := strings.Split(elizaAnswer, ";")
		resp := Response(pattern, listOfelizaAnswer)
		allResponses = append(allResponses, resp)
	} //End of for
	return allResponses

} //End of docPatterns function

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	words := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, word := range words {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], word); matched {
				words[i] = reflection[1]
				break
			}
		}
	}

	// Put the tokens back together.
	return strings.Join(words, ``)
}
func randomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(answers))
	return answers[i]
}
