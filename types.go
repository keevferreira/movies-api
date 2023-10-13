package main

import "time"

type Movie struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	ReleaseYear   int       `json:"release_year"`
	Director      string    `json:"director"`
	LeadActors    []string  `json:"lead_actors"`
	Genre         string    `json:"genre"`
	AgeRating     string    `json:"age_rating"`
	Rating        float64   `json:"rating"`
	NumVotes      int       `json:"num_votes"`
	CoverImageURL string    `json:"cover_image_url"`
	CreationDate  time.Time `json:"creation_date"`
	UpdateDate    time.Time `json:"update_date"`
}

func NewMovie(title, description, genre, director, ageRating, coverImageURL string, releaseYear int, leadActors []string) *Movie {
	return &Movie{
		Title:         title,
		Description:   description,
		Genre:         genre,
		Director:      director,
		AgeRating:     ageRating,
		CoverImageURL: coverImageURL,
		ReleaseYear:   releaseYear,
		LeadActors:    leadActors,
		CreationDate:  time.Now(),
		UpdateDate:    time.Now(),
	}
}
