package schema_test

import (
	"testing"

	"github.com/gocipe/schemaorg/schema"
)

func TestWebSite_String(t *testing.T) {
	type fields struct {
		MetaContext string
		MetaType    string
		Name        string
		URL         string
		Issn        string
		Creator     schema.Organization
		InLanguage  string
		Keywords    string
		License     schema.License
		Audience    schema.Audience
	}

	case1 := schema.NewWebSite("Culture", "https://culture.lexpress.mu/", "FR", "Culture, Authenticite, Mauricien")

	case2 := schema.WebSite{
		MetaContext: schema.MetaContext,
		MetaType:    schema.MetaWebSite,
		Name:        "Culture",
		URL:         "https://test.test.mu/",
		Issn:        "9999-9999",
		Creator: schema.Organization{
			MetaType:  schema.MetaOrganization,
			Address:   "Example Road, Example City, Example",
			Email:     "test@gmail.com",
			Telephone: "999-9999",
			Name:      "Test Case Ltd",
		},
		InLanguage: "EN, FR, KR",
		Keywords:   "Testing, Test Driven, Test Cases",
		License: schema.License{
			MetaType:      schema.MetaLicense,
			CopyrightYear: 2018,
		},
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext: schema.MetaContext,
				MetaType:    schema.MetaWebSite,
				Name:        "Culture",
				URL:         "https://culture.lexpress.mu/",
				InLanguage:  "FR",
				Keywords:    "Culture, Authenticite, Mauricien",
			},
			want: case1.String(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				MetaContext: schema.MetaContext,
				MetaType:    schema.MetaWebSite,
				Name:        "Culture",
				URL:         "https://test.test.mu/",
				Issn:        "9999-9999",
				Creator: schema.Organization{
					MetaType:  schema.MetaOrganization,
					Address:   "Example Road, Example City, Example",
					Email:     "test@gmail.com",
					Telephone: "999-9999",
					Name:      "Test Case Ltd",
				},
				InLanguage: "EN, FR, KR",
				Keywords:   "Testing, Test Driven, Test Cases",
				License: schema.License{
					MetaType:      schema.MetaLicense,
					CopyrightYear: 2018,
				},
			},
			want: case2.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := schema.WebSite{
				MetaContext: tt.fields.MetaContext,
				MetaType:    tt.fields.MetaType,
				Name:        tt.fields.Name,
				URL:         tt.fields.URL,
				Issn:        tt.fields.Issn,
				Creator:     tt.fields.Creator,
				InLanguage:  tt.fields.InLanguage,
				Keywords:    tt.fields.Keywords,
				License:     tt.fields.License,
				Audience:    tt.fields.Audience,
			}
			if got := ws.String(); got != tt.want {
				t.Errorf("WebSite.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
