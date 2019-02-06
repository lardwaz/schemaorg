package schema

//WebPage is a web page
type WebPage struct {
	MetaType    string                 `json:"@type"`
	MetaContext string                 `json:"@context"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Keywords    string                 `json:"keywords"`
	Breadcrumb  string                 `json:"breadcrumb"`
	Image       string                 `json:"image"`
	URL         string                 `json:"url"`
	Speakable   SpeakableSpecification `json:"urlspeakable"`
}
