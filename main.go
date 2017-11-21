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
	//Code adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ask", HandleAsk)
	http.ListenAndServe(":8081", nil)
}

//input := "My name is Kamila"
//input2 := "Hello how are you?"

//fmt.Println(eliza.Ask(input))
//fmt.Println(eliza.Ask(input2))

//} //End of function main

func HandleAsk(w http.ResponseWriter, r *http.Request) {
	//Code adapted from: https://siongui.github.io/2017/03/24/go-get-url-query-string-in-http-handler/
	userInput := r.URL.Query().Get("input") // extracts hello
	reply := eliza.Ask(userInput)
	fmt.Fprintf(w, reply)

}
