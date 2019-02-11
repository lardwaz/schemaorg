package schema

import "testing"

func TestWebPage_String(t *testing.T) {
	type fields struct {
		MetaType    string
		MetaContext string
		Name        string
		Description string
		Keywords    string
		Image       string
		URL         string
		Speakable   SpeakableSpecification
	}

	case1 := NewWebPage("GeoCoordinates", "https://schema.org/GeoCoordinates", "https://www.worldatlas.com/r/w728-h425-c728x425/upload/83/db/8f/geography.jpg", "Geography, Countries, Coordinates")

	case2 := WebPage{
		MetaType:    MetaWebPage,
		MetaContext: context,
		Name:        "What Is Geography?",
		Description: "Covers the essentials of Geography, one of the most fascinating and relevant areas of study.",
		Keywords:    "WorldAtlas",
		Image:       "https://www.worldatlas.com/r/w1200-h630-c1200x630/upload/83/db/8f/geography.jpg",
		URL:         "https://www.worldatlas.com/geography.html",
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaType:    MetaWebPage,
				MetaContext: context,
				Name:        "GeoCoordinates",
				Keywords:    "Geography, Countries, Coordinates",
				Image:       "https://www.worldatlas.com/r/w728-h425-c728x425/upload/83/db/8f/geography.jpg",
				URL:         "https://schema.org/GeoCoordinates",
			},
			want: case1.String(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				MetaType:    MetaWebPage,
				MetaContext: context,
				Name:        "What Is Geography?",
				Description: "Covers the essentials of Geography, one of the most fascinating and relevant areas of study.",
				Keywords:    "WorldAtlas",
				Image:       "https://www.worldatlas.com/r/w1200-h630-c1200x630/upload/83/db/8f/geography.jpg",
				URL:         "https://www.worldatlas.com/geography.html",
			},
			want: case2.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := WebPage{
				MetaType:    tt.fields.MetaType,
				MetaContext: tt.fields.MetaContext,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Keywords:    tt.fields.Keywords,
				Image:       tt.fields.Image,
				URL:         tt.fields.URL,
				Speakable:   tt.fields.Speakable,
			}
			if got := wp.String(); got != tt.want {
				t.Errorf("WebPage.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
