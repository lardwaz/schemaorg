package schema_test

import (
	"testing"

	"go.lsl.digital/lardwaz/schemaorg/schema"
)

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
		Speakable            []schema.SpeakableSpecification
		WordCount            int64
		About                schema.Thing
		Audience             schema.Audience
		Author               schema.Person
		ContentLocation      schema.Place
		ContentRating        string
		ContentReferenceTime schema.DateTime
		Contributor          schema.Organization
		CopyrightHolder      schema.Organization
		CopyrightYear        int64
		Creator              schema.Organization
		DateCreated          string
		DateModified         string
		DatePublished        string
		Editor               schema.Person
		Headline             string
		Keywords             string
		Name                 string
		URL                  string
		Image                string
	}

	article := schema.Article{
		MetaContext:    schema.MetaContext,
		MetaType:       schema.MetaArticle,
		ArticleBody:    "The indefinite article takes two forms. It’s the word a when it precedes a word that begins with a consonant.",
		ArticleSection: "The Indefinite Article",
		Backstory:      "Articles are words that define a noun as specific or unspecific. ",
		PageEnd:        "",
		PageStart:      "",
		Pagination:     "",
		WordCount:      320,
		About: schema.Thing{
			MetaType:    schema.MetaThing,
			Description: "The indefinite article indicates that a noun refers to a general idea rather than a particular thing. ",
		},
		Audience: schema.Audience{
			MetaType:     schema.MetaAudience,
			AudienceType: "Student",
			Description:  "HSC Students",
		},
		Author: schema.Person{
			MetaType:   schema.MetaPerson,
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
	article1 := schema.NewArticle("Unit Testing", "When building a great application, it is a good idea to plan for proper testing of each of the components.", "James", "Smith", "https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/", "https://testimage.com/test.jpeg")

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext:    schema.MetaContext,
				MetaType:       schema.MetaArticle,
				ArticleBody:    "The indefinite article takes two forms. It’s the word a when it precedes a word that begins with a consonant.",
				ArticleSection: "The Indefinite Article",
				Backstory:      "Articles are words that define a noun as specific or unspecific. ",
				PageEnd:        "",
				PageStart:      "",
				Pagination:     "",
				WordCount:      320,
				About: schema.Thing{
					MetaType:    "Thing",
					Description: "The indefinite article indicates that a noun refers to a general idea rather than a particular thing. ",
				},
				Audience: schema.Audience{
					MetaType:     schema.MetaAudience,
					AudienceType: "Student",
					Description:  "HSC Students",
				},
				Author: schema.Person{
					MetaType:   schema.MetaPerson,
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
			want: article.String(),
		},

		{
			name: "Test Case 2",
			fields: fields{
				Headline:  "Unit Testing",
				Name:      "Unit Testing",
				Backstory: "When building a great application, it is a good idea to plan for proper testing of each of the components.",
				Author: schema.Person{
					MetaType:   schema.MetaPerson,
					GivenName:  "James",
					FamilyName: "Smith",
				},
				URL:   "https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/",
				Image: "https://testimage.com/test.jpeg",
			},
			want: article1.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := schema.Article{
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
