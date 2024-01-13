package Handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Data struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationDates []Locations 
}

type Locations struct {
	Country string 
	City string
	Dates []string
}
type RelationsData struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func fetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return err
	}

	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		return err
	}

	return nil
}

func FetchArtistData() ([]Data, error) {
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	relationURL := "https://groupietrackers.herokuapp.com/api/relation"

	var artists []Artist
	var relation RelationsData

	err := fetchData(artistsURL, &artists)
	if err != nil {
		return nil, err
	}

	var artistsFiltered []Artist
	for _, artist := range artists {
		if artist.ID != 21 {
			artistsFiltered = append(artistsFiltered, artist)
		}
	}

	err = fetchData(relationURL, &relation)
	if err != nil {
		return nil, err
	}

	data := []Data{}
	// Update artists with dates and locations
	for i, rel := range relation.Index {
		if len(rel.DatesLocations) > 0 {
			LocationData := []Locations{}
			for location, dates:= range rel.DatesLocations {
				L:= strings.Split(location, "-")
				city := L[0]
				country := L[1]
				Location := Locations{Country: country, City: city, Dates: dates}
				LocationData = append(LocationData, Location)
			}
			if i < len(artistsFiltered) {
				ArtistData := Data{
					ID:            artistsFiltered[i].ID,
					Image:         artistsFiltered[i].Image,
					Name:          artistsFiltered[i].Name,
					Members:       artistsFiltered[i].Members,
					CreationDate:  artistsFiltered[i].CreationDate,
					FirstAlbum:    artistsFiltered[i].FirstAlbum,
					LocationDates: LocationData,
				}
				data = append(data, ArtistData)
			}
		}
	}	
	return data, nil
}