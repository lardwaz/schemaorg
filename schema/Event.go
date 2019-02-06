package schema

//Event is: an event happening at a certain time and location, such as a concert, lecture, or festival.
type Event struct {
	MetaContext               string          `json:"MetaContext"`
	MetaType                  string          `json:"@type"`
	About                     Thing           `json:"about"`
	Actor                     Person          `json:"actor"`
	AggregateRating           AggregateRating `json:"aggregateRating"`
	Attendee                  string          `json:"attendee"`
	Audience                  Audience        `json:"audience"`
	Composer                  Organization    `json:"composer"`
	Contributor               Organization    `json:"contributor"`
	Director                  Person          `json:"director"`
	DoorTime                  string          `json:"doorTime"`
	Duration                  string          `json:"duration"`
	EndDate                   string          `json:"endDate"`
	EventStatus               string          `json:"eventStatus"`
	Funder                    Organization    `json:"funder"`
	InLanguage                string          `json:"inLanguage"`
	IsAccessibleForFree       bool            `json:"isAccessibleForFree"`
	Location                  PostalAddress   `json:"location"`
	MaximumAttendeeCapacity   int64           `json:"maximumAttendeeCapacity"`
	Organizer                 Organization    `json:"organizer"`
	Performer                 Organization    `json:"performer"`
	PreviousStartDate         string          `json:"previousStartDate"`
	RecordedIn                CreativeWork    `json:"recordedIn"`
	RemainingAttendeeCapacity int64           `json:"remainingAttendeeCapacity"`
	Review                    Review          `json:"review"`
	Sponsor                   Organization    `json:"sponsor"`
	StartDate                 string          `json:"startDate"`
	Translator                Organization    `json:"translator"`
	TypicalAgeRange           string          `json:"typicalAgeRange"`
	WorkFeatured              CreativeWork    `json:"workFeatured"`
	WorkPerformed             CreativeWork    `json:"workPerformed"`
}
