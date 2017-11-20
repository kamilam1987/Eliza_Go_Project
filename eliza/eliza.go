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

//Answers for Introductions to user
var greetings = []string{
	"Hi. How are you today?",
	"hello. How are you doing?",
	"How do you do. Please tell me your problem.",
	"Hello. How are you feeling today?",
	"Hi! How are you doing?",
	"Please tell me what's been bothering you.",
	"Is something troubling you?",
}

// Ansvers for exiting .
var byes = []string{
	"Goodbye. It was nice talking to you.",
	"Thank you for talking with me.",
	"Thank you, that will be $150. Have a good day!",
	"Goodbye. This was really a nice talk.",
	"Goodbye. I'm looking forward to our next session.",
	"This was a good session, wasn't it – but time is over now. Goodbye.",
	"Maybe we could discuss this over more in our next session? Goodbye.",
	"Good-bye.",
}

//Array of greeting from user Input
var elizaQuit = []string{
	"goodbye",
	"bye",
	"quit",
	"exit",
}

//Array of byes from user Input
var elizaStarts = []string{
	"hi",
	"hello",
	"good morning",
	"good evening",
	"good afternoon",
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

//Returns random answer from array of greetings
func elizaGreetings() string {
	return randChoice(greetings)
}

// Returns random goodbye
func elizaBye() string {
	return randChoice(byes)
}

//Function Ask compares user input if matches witch patterns
func Ask(userInput string) string {
	userInput = preprocess(userInput)

	allPatterns := docPatterns()

	if IsElizaStart(userInput) {
		return elizaGreetings()
	}
	if IsQuitStatement(userInput) {
		return elizaBye()
	}
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
		elizaPattern := Response(pattern, listOfelizaAnswer)
		allResponses = append(allResponses, elizaPattern)

	} //End of for
	return allResponses

} //End of docPatterns function

func Reflect(userInput string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	words := boundaries.Split(userInput, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`i`, `you`},
		{`I`, `You`},
		{`am`, `are`},
		{`was`, `were`},
		{`i'd`, `you would`},
		{`i've`, `you have`},
		{`i'll`, `you will`},
		{`are`, `am`},
		{`you've`, `I have`},
		{`you'll`, `I will`},
		{`yours`, `mine`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	// Loop through each token, reflecting it if there's a match.
	//Code adapted from : https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
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
} //End of Reflect function

//IsElizaStart returns if the statement is a start statement
func IsElizaStart(userInput string) bool {
	userInput = preprocess(userInput)
	for _, startStatement := range elizaStarts {
		if userInput == startStatement {
			return true
		}
	}
	return false
}

// IsQuitStatement returns if the statement is a quit statement
func IsQuitStatement(userInput string) bool {
	userInput = preprocess(userInput)
	for _, quitStatement := range elizaQuit {
		if userInput == quitStatement {
			return true
		}
	}
	return false
}

func preprocess(userInput string) string {
	//Code adapted from: https://golang.org/pkg/strings/#example_Trim
	userInput = strings.TrimRight(userInput, "\n.!")
	//Code adapted from : https://golang.org/pkg/strings/#example_ToLower
	userInput = strings.ToLower(userInput)
	return userInput
}

//Function randomAnswer gets random  answers from patterns
func randomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(answers))
	return answers[i]
}

//Function randChoice gets random answers for goodbyes and introductions
func randChoice(list []string) string {
	randIndex := rand.Intn(len(list))
	return list[randIndex]
}
