package schemaorg

import "testing"

func TestWebpageMeta_String(t *testing.T) {
	type fields struct {
		MetaContext string
		MetaType    string
		Title       string
		URL         string
		Description string
		Image       string
		Keywords    string
	}

	case1 := NewWebPageMeta("List of Recipes", "http://localhost:8080/recipes",
		"All list of recipes of categories french is listed here.", "http://localhost:8080/food.jpeg")

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test Case 1",
			fields: fields{},
			want:   case1.String(),
		},
		{
			name:   "Test Case 1",
			fields: fields{},
			want:   case1.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WebpageMeta{
				MetaContext: tt.fields.MetaContext,
				MetaType:    tt.fields.MetaType,
				Title:       tt.fields.Title,
				URL:         tt.fields.URL,
				Description: tt.fields.Description,
				Image:       tt.fields.Image,
				Keywords:    tt.fields.Keywords,
			}
			if got := w.String(); got != tt.want {
				t.Errorf("WebpageMeta.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
