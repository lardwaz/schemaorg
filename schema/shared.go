package schema

//Organization such as a school, NGO, corporation, club, etc.
type Organization struct {
	MetaType     string       `json:"@type"`
	Address      string       `json:"address"`
	ContactPoint ContactPoint `json:"contactPoint"`
	Email        string       `json:"email"`
	Telephone    string       `json:"telephone"`
	Description  string       `json:"description"`
	URL          string       `json:"url"`
	Image        string       `json:"image"`
}

//ContactPoint is a way to have contact to a Person, Organization, place etc
type ContactPoint struct {
	MetaType       string `json:"@type"`
	ContactType    string `json:"contactType"`
	FaxNumber      string `json:"faxNumber"`
	HoursAvailable string `json:"hoursAvailable"`
	Email          string `json:"email"`
	Telephone      string `json:"telephone"`
}

//Thing is the most generic type of item.
type Thing struct {
	MetaType             string `json:"@type"`
	AlternateName        string `json:"alternateName"`
	Description          string `json:"description"`
	AccessMode           string `json:"accessMode"`
	AccessModeSufficient string `json:"accessModeSufficient"`
	AccessibilityFeature string `json:"accessibilityFeature"`
	AccessibilityHazard  string `json:"accessibilityHazard"`
	AccessibilitySummary string `json:"accessibilitySummary"`
	AccountablePerson    Person `json:"accountablePerson"`
}

//Person Specifies that is legally accountable.
type Person struct {
	MetaType    string       `json:"@type"`
	ID          string       `json:"ID"`
	Affiliation Organization `json:"affiliation"`
	GivenName   string       `json:"givenName"`
	FamilyName  string       `json:"FamilyName"`
	Address     string       `json:"address"`
	Email       string       `json:"email"`
	Gender      string       `json:"gender"`
	JobTitle    string       `json:"jobTitle"`
}

//OpeningHoursSpecification is a structured value providing information about the opening hours of a place or a certain service inside a place.
type OpeningHoursSpecification struct {
	MetaType  string `json:"@type"`
	Closes    string `json:"closes"`
	DayOfWeek string `json:"dayOfWeek"`
	Opens     string `json:"opens"`
}

//Audience is i.e. a group for whom something was created.
type Audience struct {
	MetaType       string             `json:"@type"`
	AudienceType   string             `json:"audienceType"`
	Description    string             `json:"description"`
	GeographicArea AdministrativeArea `json:"geographicArea"`
}

//AdministrativeArea is a geographical region, typically under the jurisdiction of a particular government.
type AdministrativeArea struct {
	MetaType string         `json:"@type"`
	Address  string         `json:"address"`
	Geo      GeoCoordinates `json:"geo"`
	HasMap   Map            `json:"hasMap"`
}

//Map is a map
type Map struct {
	MetaType string `json:"@type"`
	MapType  string `json:"mapType"`
	URL      string `json:"URL"`
}

//Place are entities that have a somewhat fixed, physical extension.
type Place struct {
	MetaType string         `json:"@type"`
	Address  PostalAddress  `json:"address"`
	Geo      GeoCoordinates `json:"geo"`
}

//GeoCoordinates is the geographic coordinates of a place or event.
type GeoCoordinates struct {
	MetaType  string `json:"@type"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

//PostalAddress is a mailing address
type PostalAddress struct {
	MetaType            string `json:"@type"`
	AddressCountry      string `json:"addressCountry"`
	AddressLocality     string `json:"addressLocality"`
	AddressRegion       string `json:"addressRegion"`
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber"`
	PostalCode          string `json:"postalCode"`
	StreetAddress       string `json:"streetAddress"`
}

//DateTime is a combination of date and time of day in the form [-]CCYY-MM-DDThh:mm:ss[Z|(+|-)hh:mm]
type DateTime struct {
	MetaType     string `json:"@type"`
	DateCreated  string `json:"dateCreated"`
	DateDeleted  string `json:"dateDeleted"`
	DateModified string `json:"dateModified"`
}

//VideoObject is a video file.
type VideoObject struct {
	MetaType       string      `json:"@type"`
	Caption        string      `json:"caption"`
	Thumbnail      string      `json:"thumbnail"`
	Transcript     string      `json:"transcript"`
	VideoFrameSize string      `json:"videoFrameSize"`
	VideoQuality   string      `json:"videoQuality"`
	Media          MediaObject `json:"media"`
}

//Review defn: A review of an item - for example, of a restaurant, movie, or store.
type Review struct {
	MetaType     string `json:"@type"`
	ItemReviewed Thing  `itemReviewed:"itemReviewed"`
	ReviewAspect string `reviewAspect:"reviewAspect"`
	ReviewBody   string `reviewBody:"reviewBody"`
	ReviewRating Rating `reviewRating:"reviewRating"`
}

//Rating is an evaluation on a numeric scale, such as 1 to 5 stars.
type Rating struct {
	MetaType     string `json:"@type"`
	BestRating   string `json:"bestRating"`
	RatingValue  string `json:"ratingValue"`
	WorstRating  string `json:"worstRating"`
	ReviewAspect string `json:"reviewAspect"`
}

//AggregateRating is the average rating based on multiple ratings or reviews.
type AggregateRating struct {
	MetaType     string `json:"@type"`
	RatingCount  int64  `json:"ratingCount"`
	ReviewCount  int64  `json:"reviewCount"`
	BestRating   string `json:"bestRating"`
	RatingValue  string `json:"ratingValue"`
	WorstRating  string `json:"worstRating"`
	ReviewAspect string `json:"reviewAspect"`
}
