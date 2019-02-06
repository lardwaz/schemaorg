package schema

//Blog is an online diary or journal located on a website.
type Blog struct {
	MetaContext string      `json:"MetaContext"`
	MetaType    string      `json:"@type"`
	BlogPost    BlogPosting `json:"blogPost"`
	Issn        string      `json:"issn"`
}

//BlogPosting is a blog post
type BlogPosting struct {
	MetaType       string                 `json:"@type"`
	SharedContent  CreativeWork           `json:"sharedContent"`
	ArticleBody    string                 `json:"articleBody"`
	ArticleSection string                 `json:"articleSection"`
	Backstory      string                 `json:"backstory"`
	PageEnd        string                 `json:"pageEnd"`
	PageStart      string                 `json:"pageStart"`
	Pagination     string                 `json:"pagination"`
	Speakable      SpeakableSpecification `json:"speakable"`
}

//CreativeWork is the most generic kind of creative work, including books, movies, photographs, software programs, etc.
type CreativeWork struct {
	MetaType        string       `json:"@type"`
	Editor          Person       `json:"editor"`
	DateCreated     string       `json:"dateCreated"`
	DateModified    string       `json:"dateModified"`
	DatePublished   string       `json:"datePublished"`
	Keywords        string       `json:"keywords"`
	PublisherPerson Person       `json:"publisherPerson"`
	PublisherOrg    Organization `json:"publisherOrg"`
	Image           ImageObject  `json:"image"`
	Video           VideoObject  `json:"video"`
}

//ImageObject is an image file
type ImageObject struct {
	MetaType             string      `json:"@type"`
	Caption              string      `json:"caption"`
	ExifData             string      `json:"exifData"`
	Thumbnail            string      `json:"thumbnail"`
	RepresentativeOfPage string      `json:"representativeOfPage"`
	Media                MediaObject `json:"media"`
}

//MediaObject is a media object, such as an image, video, or audio object embedded in a web page or a downloadable dataset
type MediaObject struct {
	MetaType          string       `json:"@type"`
	Bitrate           string       `json:"bitrate"`
	ContentSize       string       `json:"contentSize"`
	Duration          string       `json:"duration"`
	EmbedURL          string       `json:"embedURL"`
	EncodingFormat    string       `json:"encodingFormat"`
	ProductionCompany Organization `json:"productionCompany"`
	PlayerType        string       `json:"playerType"`
	UploadDate        string       `json:"uploadDate"`
	Width             string       `json:"width"`
}
