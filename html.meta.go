package schemaorg

// HTMLMeta : structure defining various params in html base template
type HTMLMeta struct {
	BaseURL      string
	OG           OgParams
	SchemaJSONld string
}

// OgParams : structure defining parameters for social media
type OgParams struct {
	Title       string
	URL         string
	Description string
	Image       string
	Keywords    string
}

//NewOgParams returns an instance of Og params
func NewOgParams(title, url, description, image, keywords string) OgParams {
	return OgParams{
		Title:       title,
		URL:         url,
		Description: description,
		Image:       image,
		Keywords:    keywords,
	}
}
