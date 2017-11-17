//Author: Kamila Michel
//Eliza project
//This project contains a chatbot web application in Go. User will be able to visit the web app throught their browser, type in some sentences and web app will reply to it.

package main

import (
	//"net/http"
	"fmt"

	"./eliza"
)

func main() {
	//code adapted from:https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	// FileServe serves index.html from static folder
	//http.Handle function tells the http package to handle all requests from static folder (index.html)
	//http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.ListenAndServe(":8088", nil)

	question := "How are you"

	fmt.Println(eliza.Ask(question))

} //End of function main
