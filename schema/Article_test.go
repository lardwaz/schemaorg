package schema

import "testing"

func TestArticle_String(t *testing.T) {
	type fields struct {
		MetaContext          string
		MetaType             string
		ArticleBody          string
		ArticleSection       string
		Backstory            string
		PageEnd              string
		PageStart            string
		Pagination           string
		Speakable            []SpeakableSpecification
		WordCount            int64
		About                Thing
		Audience             Audience
		Author               Person
		ContentLocation      Place
		ContentRating        string
		ContentReferenceTime DateTime
		Contributor          Organization
		CopyrightHolder      Organization
		CopyrightYear        int64
		Creator              Organization
		DateCreated          string
		DateModified         string
		DatePublished        string
		Editor               Person
		Headline             string
		Keywords             string
		Name                 string
		URL                  string
		Image                string
	}

	schema := Article{
		MetaContext:    context,
		MetaType:       "Article",
		ArticleBody:    "The indefinite article takes two forms. It’s the word a when it precedes a word that begins with a consonant.",
		ArticleSection: "The Indefinite Article",
		Backstory:      "Articles are words that define a noun as specific or unspecific. ",
		PageEnd:        "",
		PageStart:      "",
		Pagination:     "",
		WordCount:      320,
		About: Thing{
			MetaType:    "Thing",
			Description: "The indefinite article indicates that a noun refers to a general idea rather than a particular thing. ",
		},
		Audience: Audience{
			MetaType:     "Audience",
			AudienceType: "Student",
			Description:  "HSC Students",
		},
		Author: Person{
			MetaType:   "Person",
			GivenName:  "Jean",
			FamilyName: "Paul",
		},
		CopyrightYear: 2019,
		DateCreated:   "2019-02-01",
		DateModified:  "2019-02-02",
		DatePublished: "2019-02-05",
		Headline:      "The Return of Article",
		Keywords:      "Article, Grammar, English",
		Name:          "The Return of Article",
		URL:           "https://www.grammarly.com/blog/articles/",
		Image:         "http://i64.tinypic.com/2ob4i.jpg",
	}

	//Default Case testing
	schema1 := NewArticle("Unit Testing", "When building a great application, it is a good idea to plan for proper testing of each of the components.", "James", "Smith", "https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/", "https://testimage.com/test.jpeg")

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext:    context,
				MetaType:       "Article",
				ArticleBody:    "The indefinite article takes two forms. It’s the word a when it precedes a word that begins with a consonant.",
				ArticleSection: "The Indefinite Article",
				Backstory:      "Articles are words that define a noun as specific or unspecific. ",
				PageEnd:        "",
				PageStart:      "",
				Pagination:     "",
				WordCount:      320,
				About: Thing{
					MetaType:    "Thing",
					Description: "The indefinite article indicates that a noun refers to a general idea rather than a particular thing. ",
				},
				Audience: Audience{
					MetaType:     "Audience",
					AudienceType: "Student",
					Description:  "HSC Students",
				},
				Author: Person{
					MetaType:   "Person",
					GivenName:  "Jean",
					FamilyName: "Paul",
				},
				CopyrightYear: 2019,
				DateCreated:   "2019-02-01",
				DateModified:  "2019-02-02",
				DatePublished: "2019-02-05",
				Headline:      "The Return of Article",
				Keywords:      "Article, Grammar, English",
				Name:          "The Return of Article",
				URL:           "https://www.grammarly.com/blog/articles/",
				Image:         "http://i64.tinypic.com/2ob4i.jpg",
			},
			want: schema.String(),
		},

		{
			name: "Test Case 2",
			fields: fields{
				Headline:  "Unit Testing",
				Name:      "Unit Testing",
				Backstory: "When building a great application, it is a good idea to plan for proper testing of each of the components.",
				Author: Person{
					GivenName:  "James",
					FamilyName: "Smith",
				},
				URL:   "https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/",
				Image: "https://testimage.com/test.jpeg",
			},
			want: schema1.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Article{
				MetaContext:          tt.fields.MetaContext,
				MetaType:             tt.fields.MetaType,
				ArticleBody:          tt.fields.ArticleBody,
				ArticleSection:       tt.fields.ArticleSection,
				Backstory:            tt.fields.Backstory,
				PageEnd:              tt.fields.PageEnd,
				PageStart:            tt.fields.PageStart,
				Pagination:           tt.fields.Pagination,
				Speakable:            tt.fields.Speakable,
				WordCount:            tt.fields.WordCount,
				About:                tt.fields.About,
				Audience:             tt.fields.Audience,
				Author:               tt.fields.Author,
				ContentLocation:      tt.fields.ContentLocation,
				ContentRating:        tt.fields.ContentRating,
				ContentReferenceTime: tt.fields.ContentReferenceTime,
				Contributor:          tt.fields.Contributor,
				CopyrightHolder:      tt.fields.CopyrightHolder,
				CopyrightYear:        tt.fields.CopyrightYear,
				Creator:              tt.fields.Creator,
				DateCreated:          tt.fields.DateCreated,
				DateModified:         tt.fields.DateModified,
				DatePublished:        tt.fields.DatePublished,
				Editor:               tt.fields.Editor,
				Headline:             tt.fields.Headline,
				Keywords:             tt.fields.Keywords,
				Name:                 tt.fields.Name,
				URL:                  tt.fields.URL,
				Image:                tt.fields.Image,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("Article.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
