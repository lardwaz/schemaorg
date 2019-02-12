package schema

import "encoding/json"

//MetaWebPage holds meta type value
const MetaWebPage string = "WebPage"

//WebPage is a web page
type WebPage struct {
	MetaType    string                 `json:"@type"`
	MetaContext string                 `json:"@context"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Keywords    string                 `json:"keywords,omitempty"`
	Image       string                 `json:"image"`
	URL         string                 `json:"url"`
	Speakable   SpeakableSpecification `json:"urlspeakable,omitempty"`
}

//NewWebPage is a new instance of webpage
func NewWebPage(name, desc, url, image, keywords string) WebPage {
	return WebPage{
		MetaContext: MetaContext,
		MetaType:    MetaWebPage,
		Name:        name,
		Description: desc,
		URL:         url,
		Image:       image,
		Keywords:    keywords,
	}
}

func (wp WebPage) String() string {
	b, _ := json.Marshal(wp)
	return string(b)
}
