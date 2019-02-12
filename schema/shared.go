package schema

//MetaContext is schema org url
const MetaContext string = "http://schema.org"

//MetaTypeForGeo is the metatype tag for GeoCoordinates
const MetaTypeForGeo string = "GeoCoordinates"

//MetaOrganization holds the value for Organization
const MetaOrganization string = "Organization"

//MetaPlace holds the value for Place (location)
const MetaPlace string = "Place"

//MetaPerson holds meta type value
const MetaPerson string = "Person"

//MetaThing holds meta type value
const MetaThing string = "Thing"

//MetaContactPoint holds meta type value
const MetaContactPoint string = "ContactPoint"

//MetaOpeningHoursSpecification holds meta type value
const MetaOpeningHoursSpecification string = "OpeningHoursSpecification"

//MetaAudience holds meta type value
const MetaAudience string = "Audience"

//MetaAdministrativeArea holds meta type value
const MetaAdministrativeArea string = "AdministrativeArea"

//MetaGeoCoordinates holds meta type value
const MetaGeoCoordinates string = "GeoCoordinates"

//MetaPostalAddress holds meta type value
const MetaPostalAddress string = "PostalAddress"

//MetaDateTime holds meta type value
const MetaDateTime string = "DateTime"

//MetaVideoObject holds meta type value
const MetaVideoObject string = "VideoObject"

//MetaMap holds meta type value
const MetaMap string = "Map"

//MetaReview holds meta type value
const MetaReview string = "Review"

//MetaRating holds meta type value
const MetaRating string = "Rating"

//MetaAggregateRating holds meta type value
const MetaAggregateRating string = "AggregateRating"

//Organization such as a school, NGO, corporation, club, etc.
type Organization struct {
	MetaType     string       `json:"@type"`
	Address      string       `json:"address,omitempty"`
	ContactPoint ContactPoint `json:"contactPoint,omitempty"`
	Email        string       `json:"email,omitempty"`
	Telephone    string       `json:"telephone,omitempty"`
	Description  string       `json:"description,omitempty"`
	URL          string       `json:"url,omitempty"`
	Image        string       `json:"image,omitempty"`
	Name         string       `json:"name,omitempty"`
}

//ContactPoint is a way to have contact to a Person, Organization, place etc
type ContactPoint struct {
	MetaType       string `json:"@type"`
	ContactType    string `json:"contactType,omitempty"`
	FaxNumber      string `json:"faxNumber,omitempty"`
	HoursAvailable string `json:"hoursAvailable,omitempty"`
	Email          string `json:"email,omitempty"`
	Telephone      string `json:"telephone,omitempty"`
}

//Thing is the most generic type of item.
type Thing struct {
	MetaType             string `json:"@type"`
	AlternateName        string `json:"alternateName,omitempty"`
	Description          string `json:"description,omitempty"`
	AccessMode           string `json:"accessMode,omitempty"`
	AccessModeSufficient string `json:"accessModeSufficient,omitempty"`
	AccessibilityFeature string `json:"accessibilityFeature,omitempty"`
	AccessibilityHazard  string `json:"accessibilityHazard,omitempty"`
	AccessibilitySummary string `json:"accessibilitySummary,omitempty"`
	AccountablePerson    Person `json:"accountablePerson,omitempty"`
}

//Person Specifies that is legally accountable.
type Person struct {
	MetaType    string       `json:"@type"`
	ID          string       `json:"ID,omitempty"`
	Affiliation Organization `json:"affiliation,omitempty"`
	GivenName   string       `json:"givenName,omitempty"`
	FamilyName  string       `json:"FamilyName,omitempty"`
	Address     string       `json:"address,omitempty"`
	Email       string       `json:"email,omitempty"`
	Gender      string       `json:"gender,omitempty"`
	JobTitle    string       `json:"jobTitle,omitempty"`
}

//OpeningHoursSpecification is a structured value providing information about the opening hours of a place or a certain service inside a place.
type OpeningHoursSpecification struct {
	MetaType  string `json:"@type"`
	Closes    string `json:"closes,omitempty"`
	DayOfWeek string `json:"dayOfWeek,omitempty"`
	Opens     string `json:"opens,omitempty"`
}

//Audience is i.e. a group for whom something was created.
type Audience struct {
	MetaType       string             `json:"@type"`
	AudienceType   string             `json:"audienceType,omitempty"`
	Description    string             `json:"description,omitempty"`
	GeographicArea AdministrativeArea `json:"geographicArea,omitempty"`
}

//AdministrativeArea is a geographical region, typically under the jurisdiction of a particular government.
type AdministrativeArea struct {
	MetaType string         `json:"@type"`
	Address  string         `json:"address,omitempty"`
	Geo      GeoCoordinates `json:"geo,omitempty"`
	HasMap   Map            `json:"hasMap,omitempty"`
}

//Map is a map
type Map struct {
	MetaType string `json:"@type"`
	MapType  string `json:"mapType,omitempty"`
	URL      string `json:"URL,omitempty"`
}

//Place are entities that have a somewhat fixed, physical extension.
type Place struct {
	MetaType string         `json:"@type"`
	Address  string         `json:"address,omitempty"`
	Geo      GeoCoordinates `json:"geo,omitempty"`
}

//GeoCoordinates is the geographic coordinates of a place or event.
type GeoCoordinates struct {
	MetaType  string `json:"@type"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

//PostalAddress is a mailing address
type PostalAddress struct {
	MetaType            string `json:"@type"`
	AddressCountry      string `json:"addressCountry,omitempty"`
	AddressLocality     string `json:"addressLocality,omitempty"`
	AddressRegion       string `json:"addressRegion,omitempty"`
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber,omitempty"`
	PostalCode          string `json:"postalCode,omitempty"`
	StreetAddress       string `json:"streetAddress,omitempty"`
}

//DateTime is a combination of date and time of day in the form [-]CCYY-MM-DDThh:mm:ss[Z|(+|-)hh:mm]
type DateTime struct {
	MetaType     string `json:"@type"`
	DateCreated  string `json:"dateCreated,omitempty"`
	DateDeleted  string `json:"dateDeleted,omitempty"`
	DateModified string `json:"dateModified,omitempty"`
}

//VideoObject is a video file.
type VideoObject struct {
	MetaType       string      `json:"@type"`
	Caption        string      `json:"caption,omitempty"`
	Thumbnail      string      `json:"thumbnail,omitempty"`
	Transcript     string      `json:"transcript,omitempty"`
	VideoFrameSize string      `json:"videoFrameSize,omitempty"`
	VideoQuality   string      `json:"videoQuality,omitempty"`
	Media          MediaObject `json:"media,omitempty"`
	URL            string      `json:"url,omitempty"`
}

//Review defn: A review of an item - for example, of a restaurant, movie, or store.
type Review struct {
	MetaType     string `json:"@type"`
	ItemReviewed Thing  `itemReviewed:"itemReviewed,omitempty"`
	ReviewAspect string `reviewAspect:"reviewAspect,omitempty"`
	ReviewBody   string `reviewBody:"reviewBody,omitempty"`
	ReviewRating Rating `reviewRating:"reviewRating,omitempty"`
}

//Rating is an evaluation on a numeric scale, such as 1 to 5 stars.
type Rating struct {
	MetaType     string `json:"@type"`
	BestRating   string `json:"bestRating,omitempty"`
	RatingValue  string `json:"ratingValue,omitempty"`
	WorstRating  string `json:"worstRating,omitempty"`
	ReviewAspect string `json:"reviewAspect,omitempty"`
}

//AggregateRating is the average rating based on multiple ratings or reviews.
type AggregateRating struct {
	MetaType     string `json:"@type"`
	RatingCount  int64  `json:"ratingCount,omitempty"`
	ReviewCount  int64  `json:"reviewCount,omitempty"`
	BestRating   string `json:"bestRating,omitempty"`
	RatingValue  string `json:"ratingValue,omitempty"`
	WorstRating  string `json:"worstRating,omitempty"`
	ReviewAspect string `json:"reviewAspect,omitempty"`
}
