package schema

import "encoding/json"

//Event is: an event happening at a certain time and location, such as a concert, lecture, or festival.
// type Event struct {
// 	MetaContext               string          `json:"MetaContext,omitempty"`
// 	MetaType                  string          `json:"@type,omitempty"`
// 	About                     Thing           `json:"about,omitempty"`
// 	Actor                     Person          `json:"actor,omitempty"`
// 	AggregateRating           AggregateRating `json:"aggregateRating,omitempty"`
// 	Attendee                  string          `json:"attendee,omitempty"`
// 	Audience                  Audience        `json:"audience,omitempty"`
// 	Composer                  Organization    `json:"composer,omitempty"`
// 	Contributor               Organization    `json:"contributor,omitempty"`
// 	Director                  Person          `json:"director,omitempty"`
// 	DoorTime                  string          `json:"doorTime,omitempty"`
// 	Duration                  string          `json:"duration,omitempty"`
// 	EndDate                   string          `json:"endDate,omitempty"`
// 	EventStatus               string          `json:"eventStatus,omitempty"`
// 	Funder                    Organization    `json:"funder,omitempty"`
// 	InLanguage                string          `json:"inLanguage,omitempty"`
// 	IsAccessibleForFree       bool            `json:"isAccessibleForFree,omitempty"`
// 	Location                  PostalAddress   `json:"location,omitempty"`
// 	MaximumAttendeeCapacity   int64           `json:"maximumAttendeeCapacity,omitempty"`
// 	Name                      string          `json:"name"`
// 	Organizer                 Organization    `json:"organizer"`
// 	Performer                 Organization    `json:"performer,omitempty"`
// 	PreviousStartDate         string          `json:"previousStartDate,omitempty"`
// 	RecordedIn                CreativeWork    `json:"recordedIn,omitempty"`
// 	RemainingAttendeeCapacity int64           `json:"remainingAttendeeCapacity,omitempty"`
// 	Review                    Review          `json:"review,omitempty"`
// 	Sponsor                   Organization    `json:"sponsor,omitempty"`
// 	StartDate                 string          `json:"startDate,omitempty"`
// 	Translator                Organization    `json:"translator,omitempty"`
// 	TypicalAgeRange           string          `json:"typicalAgeRange,omitempty"`
// 	WorkFeatured              CreativeWork    `json:"workFeatured,omitempty"`
// 	WorkPerformed             CreativeWork    `json:"workPerformed,omitempty"`
// }

//Event represents schema.org's event
type Event struct {
	MetaContext string        `json:"@context"`
	MetaType    string        `json:"@type"`
	Name        string        `json:"name,omitempty"`
	StartDate   string        `json:"startDate,omitempty"`
	EndDate     string        `json:"endDate,omitempty"`
	Image       string        `json:"image,omitempty"`
	URL         string        `json:"url,omitempty"`
	Organizer   string        `json:"organizer,omitempty"`
	Description string        `json:"description,omitempty"`
	Location    EventLocation `json:"location,omitempty"`
}

//EventAddress from schema org
type EventAddress struct {
	MetaType      string `json:"@type"`
	StreetAddress string `json:"streetAddress"`
}

//EventLocation (custom) from schema org
type EventLocation struct {
	MetaType    string         `json:"@type"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Telephone   string         `json:"telephone"`
	URL         string         `json:"url"`
	Image       string         `json:"image"`
	Address     EventAddress   `json:"address"`
	Geoloc      GeoCoordinates `json:"geo,omitempty"`
	HasMap      Map            `json:"hasMap,omitempty"`
}

func (e *Event) String() (string, error) {
	b, err := json.Marshal(e)
	return string(b), err
}

//NewEvent returns a new instance of Event
func NewEvent(name, image, url, address, phone string) Event {
	return Event{
		MetaContext: context,
		MetaType:    "Event",
		Name:        name,
		URL:         url,
		Location: EventLocation{
			Telephone: phone,
			Address: EventAddress{
				StreetAddress: address,
			},
		},
	}
}
