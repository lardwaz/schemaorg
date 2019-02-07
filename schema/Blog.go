package schema

import "encoding/json"

//Blog is an online diary or journal located on a website.
type Blog struct {
	MetaContext string      `json:"MetaContext"`
	MetaType    string      `json:"@type"`
	BlogPost    BlogPosting `json:"blogPost"`
	Issn        string      `json:"issn,omitempty"`
	Name        string      `json:"name"`
	URL         string      `json:"url"`
	Image       string      `json:"image"`
}

//BlogPosting is a blog post
type BlogPosting struct {
	MetaType       string                 `json:"@type"`
	SharedContent  CreativeWork           `json:"sharedContent,omitempty"`
	ArticleBody    string                 `json:"articleBody,omitempty"`
	ArticleSection string                 `json:"articleSection,omitempty"`
	Backstory      string                 `json:"backstory"`
	PageEnd        string                 `json:"pageEnd,omitempty"`
	PageStart      string                 `json:"pageStart,omitempty"`
	Pagination     string                 `json:"pagination,omitempty"`
	Speakable      SpeakableSpecification `json:"speakable,omitempty"`
}

//CreativeWork is the most generic kind of creative work, including books, movies, photographs, software programs, etc.
type CreativeWork struct {
	MetaType        string       `json:"@type,omitempty"`
	Editor          Person       `json:"editor,omitempty"`
	DateCreated     string       `json:"dateCreated,omitempty"`
	DateModified    string       `json:"dateModified,omitempty"`
	DatePublished   string       `json:"datePublished,omitempty"`
	Keywords        string       `json:"keywords,omitempty"`
	PublisherPerson Person       `json:"publisherPerson,omitempty"`
	PublisherOrg    Organization `json:"publisherOrg,omitempty"`
	Image           ImageObject  `json:"image,omitempty"`
	Video           VideoObject  `json:"video,omitempty"`
}

//ImageObject is an image file
type ImageObject struct {
	MetaType             string      `json:"@type,omitempty"`
	Caption              string      `json:"caption,omitempty"`
	ExifData             string      `json:"exifData,omitempty"`
	Thumbnail            string      `json:"thumbnail,omitempty"`
	RepresentativeOfPage string      `json:"representativeOfPage,omitempty"`
	Media                MediaObject `json:"media,omitempty"`
	URL                  string      `json:"url,omitempty"`
}

//MediaObject is a media object, such as an image, video, or audio object embedded in a web page or a downloadable dataset
type MediaObject struct {
	MetaType          string       `json:"@type,omitempty"`
	Bitrate           string       `json:"bitrate,omitempty"`
	ContentSize       string       `json:"contentSize,omitempty"`
	Duration          string       `json:"duration,omitempty"`
	EmbedURL          string       `json:"embedURL,omitempty"`
	EncodingFormat    string       `json:"encodingFormat,omitempty"`
	ProductionCompany Organization `json:"productionCompany,omitempty"`
	PlayerType        string       `json:"playerType,omitempty"`
	UploadDate        string       `json:"uploadDate,omitempty"`
	Width             string       `json:"width,omitempty"`
	URL               string       `json:"url,omitempty"`
}

//NewBlog returns a new instance of Blog with compulsory inputs
func NewBlog(title, url, image, summary string) Blog {
	return Blog{
		Name:  title,
		URL:   url,
		Image: image,
		BlogPost: BlogPosting{
			Backstory: summary,
		},
	}
}

func (bg *Blog) String() (string, error) {
	b, err := json.Marshal(bg)
	return string(b), err
}
