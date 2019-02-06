package schema

//Movie is a cinema film
type Movie struct {
	MetaContext       string       `json:"MetaContext"`
	MetaType          string       `json:"@type"`
	Actor             string       `json:"actor"`
	AggregateRating   string       `json:"aggregateRating"`
	Author            string       `json:"author"`
	Director          string       `json:"director"`
	Description       string       `json:"description"`
	ContentRating     string       `json:"contentRating"`
	Comment           string       `json:"comment"`
	Duration          string       `json:"duration"`
	Image             string       `json:"image"`
	MusicByMusicGroup MusicGroup   `json:"musicByMusicGroup"`
	MusicByIndividual Person       `json:"musicByIndividual"`
	Name              string       `json:"name"`
	ProductionCompany Organization `json:"productionCompany"`
	SubtitleLanguage  string       `json:"subtitleLanguage"`
	Trailer           VideoObject  `json:"trailer"`
}

//MusicGroup A musical group, such as a band, an orchestra, or a choir. Can also be a solo musician.
type MusicGroup struct {
	MetaType string     `json:"@type"`
	Album    MusicAlbum `json:"album"`
	Genre    string     `json:"genre"`
}

//MusicAlbum is a collection of music tracks.
type MusicAlbum struct {
	MetaType string `json:"@type"`
	Duration string `json:"duration"`
	IsrcCode string `json:"isrcCode"`
}
