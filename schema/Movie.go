package schema

import (
	"encoding/json"
	"strings"
)

//Movie is a cinema film
type Movie struct {
	MetaContext      string      `json:"MetaContext"`
	MetaType         string      `json:"@type"`
	Actor            string      `json:"actor"`
	AggregateRating  string      `json:"aggregateRating,omitempty"`
	Author           string      `json:"author,omitempty"`
	Director         string      `json:"director"`
	Description      string      `json:"description"`
	ContentRating    string      `json:"contentRating,omitempty"`
	Comment          string      `json:"comment,omitempty"`
	Duration         string      `json:"duration,omitempty"`
	Image            string      `json:"image"`
	Name             string      `json:"name"`
	SubtitleLanguage string      `json:"subtitleLanguage,omitempty"`
	Trailer          VideoObject `json:"trailer,omitempty"`
}

//NewMovie returns a new instance of Movie with compulsory attributes set
func NewMovie(name, description, image string, director, actor []string) Movie {
	return Movie{
		MetaContext: context,
		MetaType:    "Movie",
		Name:        name,
		Description: description,
		Image:       image,
		Director:    strings.Join(director, ","),
		Actor:       strings.Join(actor, ","),
	}
}

func (m *Movie) String() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}
