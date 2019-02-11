package schema

import "encoding/json"

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
func NewWebPage(name, url, description, image, keywords string, cssSelector, xpath []string) WebPage {
	return WebPage{
		MetaContext: "Context",
		MetaType:    "WebPage",
		Name:        name,
		URL:         url,
		Description: description,
		Image:       image,
		Keywords:    keywords,
		Speakable: SpeakableSpecification{
			CSSSelector: cssSelector,
			Xpath:       xpath,
		},
	}
}

func (wp *WebPage) String() string {
	b, _ := json.Marshal(wp)
	return string(b)
}
