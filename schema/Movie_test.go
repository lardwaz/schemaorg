package schema

import (
	"strings"
	"testing"
)

func TestMovie_String(t *testing.T) {
	type fields struct {
		MetaContext      string
		MetaType         string
		Actor            string
		AggregateRating  string
		Author           string
		Director         string
		Description      string
		ContentRating    string
		Comment          string
		Duration         string
		Image            string
		Name             string
		SubtitleLanguage string
		URL              string
		Genre            string
		Trailer          VideoObject
	}

	directors := []string{"Kevin Donovan"}
	actors := []string{"Jackie Chan", "Christophe Beck", "‎John Debney"}

	case1 := NewMovie("The Tuxedo", "Cabbie-turned-chauffeur Jimmy Tong (Jackie Chan) learns there is really only one rule.", "https://en.wikipedia.org/wiki/The_Tuxedo", "http://www.gstatic.com/tv/thumb/v22vodart/29966/p29966_v_v8_aa.jpg", directors, actors)

	case2 := Movie{
		MetaContext:      context,
		MetaType:         "Movie",
		Actor:            "Kristen Bell, Jonathan Groff, Idina Menzel",
		Director:         "Jennifer Lee, Chris Buck",
		Description:      "When their kingdom becomes trapped in perpetual winter, fearless Anna (Kristen Bell) joins forces with mountaineer Kristoff.",
		Duration:         "100 mins",
		Image:            "https://upload.wikimedia.org/wikipedia/en/thumb/0/05/Frozen_%282013_film%29_poster.jpg/220px-Frozen_%282013_film%29_poster.jpg",
		Name:             "Frozen",
		SubtitleLanguage: "FR",
		URL:              "https://en.wikipedia.org/wiki/Frozen_(2013_film)",
		Genre:            "Fantasy, Comedy",
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext: context,
				MetaType:    "Movie",
				Name:        "The Tuxedo",
				Description: "Cabbie-turned-chauffeur Jimmy Tong (Jackie Chan) learns there is really only one rule.",
				Image:       "http://www.gstatic.com/tv/thumb/v22vodart/29966/p29966_v_v8_aa.jpg",
				URL:         "https://en.wikipedia.org/wiki/The_Tuxedo",
				Director:    strings.Join(directors, ", "),
				Actor:       strings.Join(actors, ", "),
			},
			want: case1.String(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				MetaContext:      context,
				MetaType:         "Movie",
				Actor:            "Kristen Bell, Jonathan Groff, Idina Menzel",
				Director:         "Jennifer Lee, Chris Buck",
				Description:      "When their kingdom becomes trapped in perpetual winter, fearless Anna (Kristen Bell) joins forces with mountaineer Kristoff.",
				Duration:         "100 mins",
				Image:            "https://upload.wikimedia.org/wikipedia/en/thumb/0/05/Frozen_%282013_film%29_poster.jpg/220px-Frozen_%282013_film%29_poster.jpg",
				Name:             "Frozen",
				SubtitleLanguage: "FR",
				URL:              "https://en.wikipedia.org/wiki/Frozen_(2013_film)",
				Genre:            "Fantasy, Comedy",
			},
			want: case2.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Movie{
				MetaContext:      tt.fields.MetaContext,
				MetaType:         tt.fields.MetaType,
				Actor:            tt.fields.Actor,
				AggregateRating:  tt.fields.AggregateRating,
				Author:           tt.fields.Author,
				Director:         tt.fields.Director,
				Description:      tt.fields.Description,
				ContentRating:    tt.fields.ContentRating,
				Comment:          tt.fields.Comment,
				Duration:         tt.fields.Duration,
				Image:            tt.fields.Image,
				Name:             tt.fields.Name,
				SubtitleLanguage: tt.fields.SubtitleLanguage,
				URL:              tt.fields.URL,
				Genre:            tt.fields.Genre,
				Trailer:          tt.fields.Trailer,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Movie.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
