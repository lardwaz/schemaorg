package schemaorg

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewOgParams(t *testing.T) {
	type args struct {
		title       string
		url         string
		description string
		image       string
		keywords    string
	}

	x := args{
		title:       "Test title",
		url:         "Test url",
		description: "Test description",
		image:       "Test image",
		keywords:    "Test keywords",
	}

	t.Run("Test New OG params", func(t *testing.T) {
		duplicate := args{
			title:       "Test title",
			url:         "Test url",
			description: "Test description",
			image:       "Test image",
			keywords:    "Test keywords",
		}

		if got := NewOgParams(duplicate.title, duplicate.url, duplicate.description, duplicate.image, duplicate.keywords); !reflect.DeepEqual(got, x) {
			t.Errorf("NewOgParams() = %v, want %v", got, x)
		}
	})

}

func TestNewMovieMeta(t *testing.T) {
	type args struct {
		name            string
		url             string
		image           string
		trailer         string
		contentrating   string
		aggregaterating string
		about           string
		description     string
		comment         string
		genre           []string
		actor           []string
		author          []string
		director        []string
	}

	x := args{
		name:            "Test name",
		url:             "Test url",
		image:           "Test image",
		trailer:         "Test trailer",
		contentrating:   "Test contentrating",
		aggregaterating: "Test aggregaterating",
		about:           "Test about",
		description:     "Test description",
		comment:         "Test comment",
		genre:           []string{},
		actor:           []string{},
		author:          []string{},
		director:        []string{},
	}
	t.Run("Test Movie", func(t *testing.T) {
		if got := NewMovieMeta(x.name, x.url, x.image, x.trailer, x.contentrating, x.aggregaterating, x.about, x.description, x.comment, x.genre, x.actor, x.author, x.director); !reflect.DeepEqual(got, x) {
			t.Errorf("NewMovieMeta() = %v, want %v", got, x)
		}
	})

}

func TestNewEventMeta(t *testing.T) {
	type args struct {
		name                string
		startdate           string
		enddate             string
		image               string
		url                 string
		organizer           string
		description         string
		locationname        string
		locationdescription string
		locationtelephone   string
		locationurl         string
		locationimage       string
		locationaddress     string
		lat                 string
		lng                 string
	}

	x := args{
		name:                "Test name",
		startdate:           "01/01/2019",
		enddate:             "12/02/2019",
		image:               "Test image",
		url:                 "Test url",
		organizer:           "Test organizer",
		description:         "Test description",
		locationname:        "Test locationname",
		locationdescription: "Test locationdescription",
		locationtelephone:   "Test locationtelephone",
		locationurl:         "Test locationurl",
		locationimage:       "Test locationimage",
		locationaddress:     "Test locationaddress",
		lat:                 "-10.85",
		lng:                 "22.14",
	}

	t.Run("Event Test", func(t *testing.T) {
		if got := NewEventMeta(
			x.name, x.startdate, x.enddate, x.image, x.url, x.organizer, x.description, x.locationname, x.locationdescription, x.locationtelephone, x.locationurl, x.locationimage, x.locationaddress, x.lat, x.lng,
		); !reflect.DeepEqual(got, x) {
			t.Errorf("NewEventMeta() = %v, want %v", got, x)
		}
	})

}

func TestNewWebPageMeta(t *testing.T) {
	type args struct {
		title       string
		url         string
		description string
		image       string
		keywords    string
	}

	x := args{
		title:       "Test title",
		url:         "Test url",
		description: "Test description",
		image:       "image.png",
		keywords:    "",
	}

	t.Run("WebPageMeta test", func(t *testing.T) {
		duplicate := args{
			title:       "Test title",
			url:         "Test url",
			description: "Test description",
			image:       "image.png",
			keywords:    "",
		}

		if got := NewWebPageMeta(duplicate.title, duplicate.url, duplicate.description, duplicate.image, duplicate.keywords); !reflect.DeepEqual(got, x) {
			t.Errorf("NewWebPageMeta() = %v, want %v", got, x)
		}
	})

}

func TestNewRecipeMeta(t *testing.T) {
	type args struct {
		name               string
		recipecuisine      string
		recipecategory     string
		description        string
		image              string
		video              string
		recipeyield        string
		cooktime           string
		preptime           string
		totaltime          string
		datepublished      string
		url                string
		ingredients        []string
		recipeinstructions []string
	}

	x := args{
		name:               "Test name",
		recipecuisine:      "Mauricien",
		recipecategory:     "Test recipecategory",
		description:        "Test description",
		image:              "image.png",
		video:              "vids.mpeg",
		recipeyield:        "Test recipeyield",
		cooktime:           "10 mins",
		preptime:           "10 mins",
		totaltime:          "10 mins",
		datepublished:      "15/01/2019",
		url:                "Test url",
		ingredients:        []string{"ingredients1", "ingredients2", "ingredients3"},
		recipeinstructions: []string{"recipeinstructions1", "recipeinstructions2", "recipeinstructions3"},
	}

	t.Run("Test", func(t *testing.T) {

		duplicate := args{
			name:               "Test name",
			recipecuisine:      "Mauricien",
			recipecategory:     "Test recipecategory",
			description:        "Test description",
			image:              "image.png",
			video:              "vids.mpeg",
			recipeyield:        "Test recipeyield",
			cooktime:           "10 mins",
			preptime:           "10 mins",
			totaltime:          "10 mins",
			datepublished:      "15/01/2019",
			url:                "Test url",
			ingredients:        []string{"ingredients1", "ingredients2", "ingredients3"},
			recipeinstructions: []string{"recipeinstructions1", "recipeinstructions2", "recipeinstructions3"},
		}

		if got := NewRecipeMeta(duplicate.name, duplicate.recipecuisine, duplicate.recipecategory, duplicate.description, duplicate.image, duplicate.video, duplicate.recipeyield, duplicate.cooktime, duplicate.preptime, duplicate.totaltime, duplicate.datepublished, duplicate.url, duplicate.ingredients, duplicate.recipeinstructions); !reflect.DeepEqual(got, x) {
			t.Errorf("NewRecipeMeta() = %v, want %v", got, x)
		}
	})

}

func TestNewVenueMeta(t *testing.T) {
	type args struct {
		name        string
		description string
		telephone   string
		url         string
		image       string
		street      string
		locality    string
		region      string
		lat         string
		lng         string
		gallery     []string
	}
	x := args{
		name:        "Test name",
		description: "Test description",
		telephone:   "Test telephone",
		url:         "Test url",
		image:       "Test image",
		street:      "Test street",
		locality:    "Test locality",
		region:      "Test region",
		lat:         "10.00",
		lng:         "-10.00",
		gallery:     []string{"image1.png", "image2.png", "image3.png"},
	}

	t.Run("Test", func(t *testing.T) {
		xcopy := args{
			name:        "Test name",
			description: "Test description",
			telephone:   "Test telephone",
			url:         "Test url",
			image:       "Test image",
			street:      "Test street",
			locality:    "Test locality",
			region:      "Test region",
			lat:         "10.00",
			lng:         "-10.00",
			gallery:     []string{"image1.png", "image2.png", "image3.png"},
		}

		if got := NewVenueMeta(xcopy.name, xcopy.description, xcopy.telephone, xcopy.url, xcopy.image, xcopy.street, xcopy.locality, xcopy.region, xcopy.lat, xcopy.lng, xcopy.gallery); !reflect.DeepEqual(got, x) {
			t.Errorf("NewVenueMeta() = %v, want %v", got, x)
		}
	})

}

func TestVenueMeta_String(t *testing.T) {
	type fields struct {
		MetaContext string
		MetaType    string
		Name        string
		Description string
		Telephone   string
		URL         string
		Image       string
		Gallery     []string
		Address     VenueAddress
		LatLng      Geo
		HasMap      Map
	}

	x := fields{
		MetaType:    "Place",
		MetaContext: Context,
		Name:        "Test name",
		Description: "Test description",
		Telephone:   "Test telephone",
		URL:         "Test url",
		Image:       "Test image",
		Gallery:     []string{"www.dummy.com/image.png", "www.dummy.com/image.png"},
		Address: VenueAddress{
			MetaType:        "PostalAddress",
			StreetAddress:   "Test street",
			AddressLocality: "Test locality",
			AddressRegion:   "Test region",
		},
		LatLng: Geo{
			MetaType:  "GeoCoordinates",
			Latitude:  "14.02",
			Longitude: "14.02",
		},
		HasMap: Map{
			MetaType: "Map",
			URL:      "https://www.google.com/maps/@" + "14.02" + ", " + "14.02" + " ,11z",
			Mtype: MapType{
				ID: "http://schema.org/VenueMap",
			},
		},
	}

	validmarshalled, _ := json.Marshal(x)

	y := fields{
		MetaType:    "Place",
		MetaContext: Context,
		Name:        "",
		Description: "",
		Telephone:   "",
		URL:         "",
		Image:       "",
		Gallery:     []string{},
		Address: VenueAddress{
			MetaType:        "PostalAddress",
			StreetAddress:   "",
			AddressLocality: "",
			AddressRegion:   "",
		},
		LatLng: Geo{
			MetaType:  "GeoCoordinates",
			Latitude:  "",
			Longitude: "",
		},
		HasMap: Map{
			MetaType: "Map",
			URL:      "https://www.google.com/maps/@" + "" + ", " + "" + " ,11z",
			Mtype: MapType{
				ID: "http://schema.org/VenueMap",
			},
		},
	}

	marshalled, _ := json.Marshal(y)

	t.Run("test", func(t *testing.T) {
		w := &VenueMeta{
			MetaType:    "Place",
			MetaContext: Context,
			Name:        "Test name",
			Description: "Test description",
			Telephone:   "Test telephone",
			URL:         "Test url",
			Image:       "Test image",
			Gallery:     []string{"www.dummy.com/image.png", "www.dummy.com/image.png"},
			Address: VenueAddress{
				MetaType:        "PostalAddress",
				StreetAddress:   "Test street",
				AddressLocality: "Test locality",
				AddressRegion:   "Test region",
			},
			LatLng: Geo{
				MetaType:  "GeoCoordinates",
				Latitude:  "14.02",
				Longitude: "14.02",
			},
			HasMap: Map{
				MetaType: "Map",
				URL:      "https://www.google.com/maps/@" + "14.02" + ", " + "14.02" + " ,11z",
				Mtype: MapType{
					ID: "http://schema.org/VenueMap",
				},
			},
		}
		if got, _ := w.String(); got != string(validmarshalled) {
			t.Errorf("VenueMeta.String() = %v, want %v", got, string(validmarshalled))
		}

		i := &VenueMeta{
			MetaType:    "Place",
			MetaContext: Context,
			Name:        "",
			Description: "",
			Telephone:   "",
			URL:         "",
			Image:       "",
			Gallery:     []string{},
			Address: VenueAddress{
				MetaType:        "PostalAddress",
				StreetAddress:   "",
				AddressLocality: "",
				AddressRegion:   "",
			},
			LatLng: Geo{
				MetaType:  "GeoCoordinates",
				Latitude:  "",
				Longitude: "",
			},
			HasMap: Map{
				MetaType: "Map",
				URL:      "https://www.google.com/maps/@" + "" + ", " + "" + " ,11z",
				Mtype: MapType{
					ID: "http://schema.org/VenueMap",
				},
			},
		}

		if got, _ := i.String(); got != string(marshalled) {
			t.Errorf("MovieMeta.String() = %v, want %v", got, string(marshalled))
		}
	})
}

func TestRecipeMeta_String(t *testing.T) {
	type fields struct {
		MetaContext        string
		MetaType           string
		Name               string
		RecipeCuisine      string
		RecipeCategory     string
		Description        string
		Image              string
		Video              string
		Ingredients        []string
		RecipeInstructions []string
		RecipeYield        string
		CookTime           string
		PrepTime           string
		TotalTime          string
		DatePublished      string
		URL                string
	}

	x := fields{
		MetaContext:        Context,
		MetaType:           "Recipe",
		Name:               "Test Name",
		RecipeCuisine:      "Test RecipeCuisine",
		RecipeCategory:     "Test RecipeCategory",
		Description:        "Test Description",
		Image:              "Test Image",
		Video:              "Test Video",
		Ingredients:        []string{"Egg", "Cabagges", "Chilli"},
		RecipeInstructions: []string{"Beat eggs", "Cut cabagges"},
		RecipeYield:        "",
		CookTime:           "10 mins",
		PrepTime:           "10 mins",
		TotalTime:          "10 mins",
		DatePublished:      "12/01/2019",
		URL:                "",
	}

	validmarshalled, _ := json.Marshal(x)

	y := fields{
		MetaContext:        Context,
		MetaType:           "Recipe",
		Name:               "",
		RecipeCuisine:      "",
		RecipeCategory:     "",
		Description:        "",
		Image:              "",
		Video:              "",
		Ingredients:        []string{},
		RecipeInstructions: []string{},
		RecipeYield:        "",
		CookTime:           "",
		PrepTime:           "",
		TotalTime:          "",
		DatePublished:      "",
		URL:                "",
	}

	marshalled, _ := json.Marshal(y)

	t.Run("test", func(t *testing.T) {
		w := &RecipeMeta{
			MetaContext:        Context,
			MetaType:           "Recipe",
			Name:               "",
			RecipeCuisine:      "",
			RecipeCategory:     "",
			Description:        "",
			Image:              "",
			Video:              "",
			Ingredients:        []string{},
			RecipeInstructions: []string{},
			RecipeYield:        "",
			CookTime:           "",
			PrepTime:           "",
			TotalTime:          "",
			DatePublished:      "",
			URL:                "",
		}
		if got, _ := w.String(); got != string(validmarshalled) {
			t.Errorf("RecipeMeta.String() = %v, want %v", got, string(validmarshalled))
		}

		i := &RecipeMeta{
			MetaContext:        "",
			MetaType:           "",
			Name:               "",
			RecipeCuisine:      "",
			RecipeCategory:     "",
			Description:        "",
			Image:              "",
			Video:              "",
			Ingredients:        []string{},
			RecipeInstructions: []string{},
			RecipeYield:        "",
			CookTime:           "",
			PrepTime:           "",
			TotalTime:          "",
			DatePublished:      "",
			URL:                "",
		}

		if got, _ := i.String(); got != string(marshalled) {
			t.Errorf("MovieMeta.String() = %v, want %v", got, string(marshalled))
		}
	})
}

func TestMovieMeta_String(t *testing.T) {
	type fields struct {
		MetaContext     string
		MetaType        string
		Name            string
		URL             string
		Image           string
		Trailer         string
		ContentRating   string
		AggregateRating string
		About           string
		Description     string
		Comment         string
		Genre           []string
		Actor           []Person
		Author          []Person
		Director        []Person
	}

	g := []string{""}

	x := fields{
		MetaContext:     Context,
		MetaType:        "Movie",
		Name:            "This is a test Name",
		URL:             "This is a test URL",
		Image:           "This is a test Image",
		Trailer:         "This is a test Trailer",
		ContentRating:   "This is a test ContentRating",
		AggregateRating: "This is a test AggregateRating",
		About:           "This is a test About",
		Description:     "This is a test Description",
		Comment:         "This is a test Comment",
		Genre:           []string{"Adventure", "Comedie"},
		Actor:           []Person{{Name: "Shakti K Helmes"}, {Name: "Leonardo Di Caprice"}},
		Author:          []Person{{Name: "Lisa Smith"}, {Name: "Eliana Homes"}},
		Director:        []Person{{Name: "Rahul Kumar"}, {Name: "Sam Blacksmith"}},
	}

	validmarshalled, _ := json.Marshal(x)

	y := fields{
		MetaContext:     Context,
		MetaType:        "Movie",
		Name:            "",
		URL:             "",
		Image:           "",
		Trailer:         "",
		ContentRating:   "",
		AggregateRating: "",
		About:           "",
		Description:     "",
		Comment:         "",
		Genre:           g,
		Actor:           []Person{},
		Author:          []Person{},
		Director:        []Person{},
	}

	marshalled, _ := json.Marshal(y)

	t.Run("test", func(t *testing.T) {
		w := &MovieMeta{
			MetaContext:     Context,
			MetaType:        "Movie",
			Name:            "",
			URL:             "",
			Image:           "",
			Trailer:         "",
			ContentRating:   "",
			AggregateRating: "",
			About:           "",
			Description:     "",
			Comment:         "",
			Genre:           g,
			Actor:           []Person{},
			Author:          []Person{},
			Director:        []Person{},
		}
		if got, _ := w.String(); got != string(validmarshalled) {
			t.Errorf("MovieMeta.String() = %v, want %v", got, string(validmarshalled))
		}

		i := &MovieMeta{
			MetaContext:     Context,
			MetaType:        "Movie",
			Name:            "",
			URL:             "",
			Image:           "",
			Trailer:         "",
			ContentRating:   "",
			AggregateRating: "",
			About:           "",
			Description:     "",
			Comment:         "",
			Genre:           g,
			Actor:           []Person{},
			Author:          []Person{},
			Director:        []Person{},
		}

		if got, _ := i.String(); got != string(marshalled) {
			t.Errorf("MovieMeta.String() = %v, want %v", got, string(marshalled))
		}
	})

}

func TestEventMeta_String(t *testing.T) {
	type fields struct {
		MetaContext string
		MetaType    string
		Name        string
		StartDate   string
		EndDate     string
		Image       string
		URL         string
		Organizer   string
		Description string
		Location    EventLocation
	}

	x := fields{
		MetaContext: Context,
		MetaType:    "Event",
		Name:        "Test Event",
		StartDate:   "12/01/2019",
		EndDate:     "19/01/2019",
		Image:       "www.dummyimage.com/image/dummy-image.jpeg",
		URL:         "",
		Organizer:   "Test Organizer",
		Description: "This is a test description",
		Location: EventLocation{
			MetaType:    "Place",
			Description: "This is a test location description",
			Telephone:   "This is a test location telephone",
			URL:         "This is a test location url",
			Image:       "This is a test location image",
			Address: EventAddress{
				MetaType:      "PostalAddress",
				StreetAddress: "This is a test location address",
			},
			Geoloc: Geo{
				MetaType:  "GeoCoordinates",
				Latitude:  "1.12",
				Longitude: "1.12",
			},
			HasMap: Map{
				MetaType: "Map",
				URL:      "https://www.google.com/maps/@" + "1.12" + ", " + "1.12" + " ,11z",
				Mtype: MapType{
					ID: "http://schema.org/VenueMap",
				},
			},
		},
	}

	validmarshalled, _ := json.Marshal(x)

	y := fields{
		MetaContext: Context,
		MetaType:    "Event",
		Name:        "",
		StartDate:   "",
		EndDate:     "",
		Image:       "",
		URL:         "",
		Organizer:   "",
		Description: "",
		Location: EventLocation{
			MetaType:    "Place",
			Description: "",
			Telephone:   "",
			URL:         "",
			Image:       "",
			Address: EventAddress{
				MetaType:      "PostalAddress",
				StreetAddress: "",
			},
			Geoloc: Geo{
				MetaType:  "GeoCoordinates",
				Latitude:  "",
				Longitude: "",
			},
			HasMap: Map{
				MetaType: "Map",
				URL:      "https://www.google.com/maps/@" + "" + ", " + "" + " ,11z",
				Mtype: MapType{
					ID: "http://schema.org/VenueMap",
				},
			},
		},
	}

	marshalled, _ := json.Marshal(y)

	t.Run("test", func(t *testing.T) {
		w := &EventMeta{
			MetaContext: Context,
			MetaType:    "Event",
			Name:        "Test Event",
			StartDate:   "12/01/2019",
			EndDate:     "19/01/2019",
			Image:       "www.dummyimage.com/image/dummy-image.jpeg",
			URL:         "",
			Organizer:   "Test Organizer",
			Description: "This is a test description",
			Location: EventLocation{
				MetaType:    "Place",
				Description: "This is a test location description",
				Telephone:   "This is a test location telephone",
				URL:         "This is a test location url",
				Image:       "This is a test location image",
				Address: EventAddress{
					MetaType:      "PostalAddress",
					StreetAddress: "This is a test location address",
				},
				Geoloc: Geo{
					MetaType:  "GeoCoordinates",
					Latitude:  "1.12",
					Longitude: "1.12",
				},
				HasMap: Map{
					MetaType: "Map",
					URL:      "https://www.google.com/maps/@" + "1.12" + ", " + "1.12" + " ,11z",
					Mtype: MapType{
						ID: "http://schema.org/VenueMap",
					},
				},
			},
		}
		if got, _ := w.String(); got != string(validmarshalled) {
			t.Errorf("EventMeta.String() = %v, want %v", got, string(validmarshalled))
		}

		i := &EventMeta{
			MetaContext: Context,
			MetaType:    "Event",
			Name:        "",
			StartDate:   "",
			EndDate:     "",
			Image:       "",
			URL:         "",
			Organizer:   "",
			Description: "",
			Location: EventLocation{
				MetaType:    "Place",
				Description: "",
				Telephone:   "",
				URL:         "",
				Image:       "",
				Address: EventAddress{
					MetaType:      "PostalAddress",
					StreetAddress: "",
				},
				Geoloc: Geo{
					MetaType:  "GeoCoordinates",
					Latitude:  "",
					Longitude: "",
				},
				HasMap: Map{
					MetaType: "Map",
					URL:      "https://www.google.com/maps/@" + "" + ", " + "" + " ,11z",
					Mtype: MapType{
						ID: "http://schema.org/VenueMap",
					},
				},
			},
		}

		if got, _ := i.String(); got != string(marshalled) {
			t.Errorf("EventMeta.String() = %v, want %v", got, string(marshalled))
		}
	})

}

func TestWebpageMeta_String(t *testing.T) {
	type fields struct {
		MetaContext string `json:"@context"`
		MetaType    string `json:"@type"`
		Title       string `json:"title"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Keywords    string `json:"keywords"`
	}

	x := fields{
		MetaContext: Context,
		MetaType:    "WebPage",
		Title:       "Testing 1",
		URL:         "www.testing.com",
		Description: "Lorem ipsum dolore amet it sem",
		Image:       "http://localhost:8080/image/image1",
		Keywords:    "",
	}

	validmarshalled, _ := json.Marshal(x)

	y := fields{
		MetaContext: Context,
		MetaType:    "WebPage",
		Title:       "",
		URL:         "",
		Description: "",
		Image:       "",
		Keywords:    "",
	}

	marshalled, _ := json.Marshal(y)

	t.Run("test", func(t *testing.T) {
		w := &WebpageMeta{
			MetaContext: Context,
			MetaType:    "WebPage",
			Title:       "Testing 1",
			URL:         "www.testing.com",
			Description: "Lorem ipsum dolore amet it sem",
			Image:       "http://localhost:8080/image/image1",
			Keywords:    "",
		}
		if got, _ := w.String(); got != string(validmarshalled) {
			t.Errorf("WebpageMeta.String() = %v, want %v", got, string(validmarshalled))
		}

		i := &WebpageMeta{
			MetaContext: Context,
			MetaType:    "WebPage",
			Title:       "",
			URL:         "",
			Description: "",
			Image:       "",
			Keywords:    "",
		}

		if got, _ := i.String(); got != string(marshalled) {
			t.Errorf("WebpageMeta.String() = %v, want %v", got, string(marshalled))
		}
	})

}