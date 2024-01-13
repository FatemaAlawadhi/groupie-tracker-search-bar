package main

import (
	"fmt"
	"log"
	"net/http"
	"groupie-tracker-search-bar/Handlers"
)

var err error 
func main() {
	http.HandleFunc("/", Handlers.HandleHomePage)
	http.HandleFunc("/stars", Handlers.HandleStarsPage)
	http.HandleFunc("/about", Handlers.HandleAboutPage)
	http.HandleFunc("/stardetails/", Handlers.HandleStarDetailsPage)
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fmt.Println("starting server at port 8143\n")
	err = http.ListenAndServe(":8143", nil)
	if err != nil {
		log.Fatal(err)
	}
}
