package schema

const suffixForDayOfWeek = "http://schema.org/"

//Article such as a news article or piece of investigative report. Newspapers and magazines have articles of many different types and this is intended to cover them all.
type Article struct {
	MetaContext          string                   `json:"MetaContext"`
	MetaType             string                   `json:"@type"`
	ArticleBody          string                   `json:"articleBody"`
	ArticleSection       string                   `json:"articleSection"`
	Backstory            string                   `json:"backstory"`
	PageEnd              string                   `json:"pageEnd"`
	PageStart            string                   `json:"pageStart"`
	Pagination           string                   `json:"pagination"`
	Speakable            []SpeakableSpecification `json:"speakable"`
	WordCount            int64                    `json:"wordCount"`
	About                Thing                    `json:"about"`
	Audience             Audience                 `json:"audience"`
	Author               Person                   `json:"author"`
	ContentLocation      Place                    `json:"contentLocation"`
	ContentRating        string                   `json:"contentRating"`
	ContentReferenceTime DateTime                 `json:"contentReferenceTime"`
	Contributor          Organization             `json:"contributor"`
	CopyrightHolder      Organization             `json:"copyrightHolder"`
	CopyrightYear        int64                    `json:"copyrightYear"`
	Creator              Organization             `json:"creator"`
	DateCreated          string                   `json:"dateCreated"`
	DateModified         string                   `json:"dateModified"`
	DatePublished        string                   `json:"datePublished"`
	Editor               Person                   `json:"editor"`
	Headline             string                   `json:"headline"`
	Keywords             string                   `json:"keywords"`
}

// SpeakableSpecification indicates sections of a Web page that are particularly 'speakable' in the sense of being highlighted as being especially appropriate for text-to-speech conversion.
type SpeakableSpecification struct {
	MetaType    string   `json:"@type"`
	CSSSelector []string `json:"cssSelector"`
	Xpath       []string `json:"xpath"`
}
