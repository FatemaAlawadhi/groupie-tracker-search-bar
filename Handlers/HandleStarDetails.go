package Handlers

import (
	"groupie-tracker-search-bar/Error"
	"html/template"
	"net/http"
	"strings"
	"strconv"
)


func HandleStarDetailsPage(w http.ResponseWriter, r *http.Request) {
	// Fetch the main data
	idstr := strings.TrimPrefix(r.URL.Path, "/stardetails/")
	ArtistsData, err:= FetchArtistData()
	if err != nil {
		Error.RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error sug")
		return
	}
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		Error.RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error sug")
		return
	} 
	var num int
	if idint >= 21 {
		num = idint-2
	} else {
		num = idint-1
	}
	
	ArtistData :=ArtistsData[num]


	
	tmpl, err := template.ParseFiles("WebPages/StarDetails.html")
	if err != nil {
		Error.RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	err = tmpl.Execute(w, ArtistData)
	if err != nil {
		Error.RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
	}
}
