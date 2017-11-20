//Author: Kamila Michel
//Eliza project
//This project contains a chatbot web application in Go. User will be able to visit the web app throught their browser, type in some sentences and web app will reply to it.

package main

import (
	//"net/http"
	"fmt"
	"net/http"

	"./eliza"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/ask", HandleAsk)
	http.ListenAndServe(":8088", nil)
}

//input := "My name is Kamila"
//input2 := "Hello how are you?"

//fmt.Println(eliza.Ask(input))
//fmt.Println(eliza.Ask(input2))

//} //End of function main

func HandleAsk(w http.ResponseWriter, r *http.Request) {
	// r.URL == "http:/localhost:8080/ask?input=hello"
	userInput := r.URL.Query().Get("input") // extracts hello
	// if userInput == ""{
	// don't respond
	// }
	reply := eliza.Ask(userInput)
	fmt.Fprintf(w, reply)

}
