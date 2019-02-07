package schema

import "encoding/json"

const suffixForDayOfWeek = "http://schema.org/"

//Article such as a news article or piece of investigative report. Newspapers and magazines have articles of many different types and this is intended to cover them all.
type Article struct {
	MetaContext          string                   `json:"MetaContext"`
	MetaType             string                   `json:"@type"`
	ArticleBody          string                   `json:"articleBody,omitempty"`
	ArticleSection       string                   `json:"articleSection,omitempty"`
	Backstory            string                   `json:"backstory"`
	PageEnd              string                   `json:"pageEnd,omitempty"`
	PageStart            string                   `json:"pageStart,omitempty"`
	Pagination           string                   `json:"pagination,omitempty"`
	Speakable            []SpeakableSpecification `json:"speakable,omitempty"`
	WordCount            int64                    `json:"wordCount,omitempty"`
	About                Thing                    `json:"about,omitempty"`
	Audience             Audience                 `json:"audience,omitempty"`
	Author               Person                   `json:"author,omitempty"`
	ContentLocation      Place                    `json:"contentLocation,omitempty"`
	ContentRating        string                   `json:"contentRating,omitempty"`
	ContentReferenceTime DateTime                 `json:"contentReferenceTime,omitempty"`
	Contributor          Organization             `json:"contributor,omitempty"`
	CopyrightHolder      Organization             `json:"copyrightHolder,omitempty"`
	CopyrightYear        int64                    `json:"copyrightYear,omitempty"`
	Creator              Organization             `json:"creator,omitempty"`
	DateCreated          string                   `json:"dateCreated,omitempty"`
	DateModified         string                   `json:"dateModified,omitempty"`
	DatePublished        string                   `json:"datePublished,omitempty"`
	Editor               Person                   `json:"editor,omitempty"`
	Headline             string                   `json:"headline,omitempty"`
	Keywords             string                   `json:"keywords"`
	Name                 string                   `json:"name"`
	URL                  string                   `json:"url"`
	Image                string                   `json:"image"`
}

// SpeakableSpecification indicates sections of a Web page that are particularly 'speakable' in the sense of being highlighted as being especially appropriate for text-to-speech conversion.
type SpeakableSpecification struct {
	MetaType    string   `json:"@type"`
	CSSSelector []string `json:"cssSelector,omitempty"`
	Xpath       []string `json:"xpath,omitempty"`
}

//NewArticle returns a new instance of Article with compulsory inputs
func NewArticle(headline, summary, authorFirstName, authorFamilyName, url, image string) Article {
	return Article{
		Headline:  headline,
		Name:      headline,
		Backstory: summary,
		Author: Person{
			GivenName:  authorFirstName,
			FamilyName: authorFamilyName,
		},
		URL:   url,
		Image: image,
	}
}

func (a *Article) String() (string, error) {
	b, err := json.Marshal(a)
	return string(b), err
}
