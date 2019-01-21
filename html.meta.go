package schemaorg

import (
	"encoding/json"
)

const Context = "http://schema.org"

// HTMLMeta : structure defining various params in html base template
type HTMLMeta struct {
	BaseURL      string
	OG           OgParams
	SchemaJSONld string
}

// OgParams : structure defining parameters for social media
type OgParams struct {
	Title       string
	URL         string
	Description string
	Image       string
	Keywords    string
}

//WebpageMeta represents schema.org's webpage
type WebpageMeta struct {
	MetaContext string `json:"@context"`
	MetaType    string `json:"@type"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Keywords    string `json:"keywords"`
}

//EventMeta represents schema.org's event
type EventMeta struct {
	MetaContext string        `json:"@context"`
	MetaType    string        `json:"@type"`
	Name        string        `json:"name"`
	StartDate   string        `json:"startDate"`
	EndDate     string        `json:"endDate"`
	Image       string        `json:"image"`
	URL         string        `json:"url"`
	Organizer   string        `json:"organizer"`
	Description string        `json:"description"`
	Location    EventLocation `json:"location"`
}
type EventAddress struct {
	MetaType      string `json:"@type"`
	StreetAddress string `json:"streetAddress"`
}

type Geo struct {
	MetaType  string `json:"@type"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Map struct {
	MetaType string  `json:"@type"`
	Mtype    MapType `json:"mapType"`
	URL      string  `json:"url"`
}

type MapType struct {
	ID string `json:"@id"`
}

type EventLocation struct {
	MetaType    string       `json:"@type"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Telephone   string       `json:"telephone"`
	URL         string       `json:"url"`
	Image       string       `json:"image"`
	Address     EventAddress `json:"address"`
	Geoloc      Geo          `json:"geo"`
	HasMap      Map          `json:"hasMap"`
}

type Person struct {
	MetaType string `json:"@type"`
	Name     string `json:"name"`
}

// MovieMeta represents schema.org's movie
type MovieMeta struct {
	MetaContext     string   `json:"@context"`
	MetaType        string   `json:"@type"`
	Name            string   `json:"name"`
	URL             string   `json:"url"`
	Image           string   `json:"image"`
	Trailer         string   `json:"trailer"`
	ContentRating   string   `json:"contentRating"`
	AggregateRating string   `json:"aggregateRating"`
	About           string   `json:"about"`
	Description     string   `json:"description"`
	Comment         string   `json:"comment"`
	Genre           []string `json:"genre"`
	Actor           []Person `json:"actor"`
	Author          []Person `json:"author"`
	Director        []Person `json:"director"`
}

// RecipeMeta represents schema.org's recipe. @see https://schema.org/Recipe
// Format for CookTime, PrepTime, TotalTime: https://en.wikipedia.org/wiki/ISO_8601#Durations
type RecipeMeta struct {
	MetaContext        string   `json:"@context"`
	MetaType           string   `json:"@type"`
	Name               string   `json:"name"`
	RecipeCuisine      string   `json:"datePublished"`
	RecipeCategory     string   `json:"description"`
	Description        string   `json:"url"`
	Image              string   `json:"image"`
	Video              string   `json:"video"`
	Ingredients        []string `json:"recipeCuisine"`
	RecipeInstructions []string `json:"recipeCategory"`
	RecipeYield        string   `json:"ingredients"`
	CookTime           string   `json:"recipeInstructions"`
	PrepTime           string   `json:"recipeYield"`
	TotalTime          string   `json:"cookTime"`
	DatePublished      string   `json:"prepTime"`
	URL                string   `json:"totalTime"`
}

type VenueAddress struct {
	MetaType        string `json:"@type"`
	StreetAddress   string `json:"streetAddress"`
	AddressLocality string `json:"addressLocality"`
	AddressRegion   string `json:"addressRegion"`
}

// VenueMeta represents schema.org's venue
type VenueMeta struct {
	MetaContext string       `json:"@context"`
	MetaType    string       `json:"@type"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Telephone   string       `json:"telephone"`
	URL         string       `json:"url"`
	Image       string       `json:"image"`
	Gallery     []string     `json:"photos"`
	Address     VenueAddress `json:"address"`
	LatLng      Geo          `json:"geo"`
	HasMap      Map          `json:"hasMap"`
}

//NewOgParams returns an instance of Og params
func NewOgParams(title, url, description, image, keywords string) OgParams {
	return OgParams{
		Title:       title,
		URL:         url,
		Description: description,
		Image:       image,
		Keywords:    keywords,
	}
}

//NewMovieMeta returns a movie meta
func NewMovieMeta(name, url, image, trailer, contentrating, aggregaterating, about, description, comment string, genre, actor, author, director []string) MovieMeta {

	actors := []Person{}
	for _, p := range actor {
		person := Person{}
		person.Name = p
		actors = append(actors, person)

	}

	authors := []Person{}
	for _, p := range author {
		person := Person{}
		person.Name = p
		authors = append(authors, person)

	}

	directors := []Person{}
	for _, p := range director {
		person := Person{}
		person.Name = p
		directors = append(directors, person)

	}

	return MovieMeta{
		MetaContext:     Context,
		MetaType:        "Movie",
		Name:            name,
		URL:             url,
		Image:           image,
		Trailer:         trailer,
		ContentRating:   contentrating,
		AggregateRating: aggregaterating,
		About:           about,
		Description:     description,
		Comment:         comment,
		Genre:           genre,
		Actor:           actors,
		Author:          authors,
		Director:        directors,
	}
}

//NewEventMeta returns a movie meta
func NewEventMeta(name, startdate, enddate, image, url, organizer, description, locationname, locationdescription, locationtelephone, locationurl, locationimage, locationaddress, lat, lng string) EventMeta {
	return EventMeta{
		MetaContext: Context,
		MetaType:    "Event",
		Name:        name,
		StartDate:   startdate,
		EndDate:     enddate,
		Image:       image,
		URL:         url,
		Organizer:   organizer,
		Description: description,
		Location: EventLocation{
			MetaType:    "Place",
			Description: locationdescription,
			Telephone:   locationtelephone,
			URL:         locationurl,
			Image:       locationimage,
			Address: EventAddress{
				MetaType:      "PostalAddress",
				StreetAddress: locationaddress,
			},
			Geoloc: Geo{
				MetaType:  "GeoCoordinates",
				Latitude:  lat,
				Longitude: lng,
			},
			HasMap: Map{
				MetaType: "Map",
				Mtype: MapType{
					ID: "http://schema.org/VenueMap",
				},
				URL: "https://www.google.com/maps/@" + lat + "," + lng + ",11z",
			},
		},
	}
}

//NewWebPageMeta returns a movie meta
func NewWebPageMeta(title, url, description, image, keywords string) WebpageMeta {
	return WebpageMeta{
		MetaContext: Context,
		MetaType:    "WebPage",
		Title:       title,
		URL:         url,
		Description: description,
		Image:       image,
		Keywords:    keywords,
	}
}

//NewRecipeMeta returns a movie meta
func NewRecipeMeta(name, recipecuisine, recipecategory, description, image, video, recipeyield, cooktime, preptime, totaltime, datepublished, url string, ingredients, recipeinstructions []string) RecipeMeta {
	return RecipeMeta{
		MetaContext:        Context,
		MetaType:           "Recipe",
		Name:               name,
		RecipeCuisine:      recipecuisine,
		RecipeCategory:     recipecategory,
		Description:        description,
		Image:              image,
		Video:              video,
		Ingredients:        ingredients,
		RecipeInstructions: recipeinstructions,
		RecipeYield:        recipeyield,
		CookTime:           cooktime,
		PrepTime:           preptime,
		TotalTime:          totaltime,
		DatePublished:      datepublished,
		URL:                url,
	}
}

//NewVenueMeta returns a movie meta
func NewVenueMeta(name, description, telephone, url, image, street, locality, region, lat, lng string, gallery []string) VenueMeta {
	return VenueMeta{
		MetaType:    "Place",
		MetaContext: Context,
		Name:        name,
		Description: description,
		Telephone:   telephone,
		URL:         url,
		Image:       image,
		Gallery:     gallery,
		Address: VenueAddress{
			MetaType:        "PostalAddress",
			StreetAddress:   street,
			AddressLocality: locality,
			AddressRegion:   region,
		},
		LatLng: Geo{
			MetaType:  "GeoCoordinates",
			Latitude:  lat,
			Longitude: lng,
		},
		HasMap: Map{
			MetaType: "Map",
			Mtype:    MapType{ID: "http://schema.org/VenueMap"},
			URL:      "https://www.google.com/maps/@" + lat + ", " + lng + " ,11z",
		},
	}
}

func (v *VenueMeta) String() (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

func (r *RecipeMeta) String() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

func (m *MovieMeta) String() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (e *EventMeta) String() (string, error) {
	b, err := json.Marshal(e)
	return string(b), err
}

func (w *WebpageMeta) String() (string, error) {
	b, err := json.Marshal(w)
	return string(b), err
}
