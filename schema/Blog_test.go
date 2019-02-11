package schema

import "testing"

func TestBlog_String(t *testing.T) {
	type fields struct {
		MetaContext string
		MetaType    string
		BlogPost    BlogPosting
		Issn        string
		Name        string
		URL         string
		Image       string
	}

	blog := NewBlog("Unit Testing In Go", "https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/", "https://testimage.com/test.jpeg", "To create basic tests for functions as well as basic tests for HTTP endpoints.")

	blog1 := Blog{
		MetaContext: context,
		MetaType:    "Blog",
		Name:        "How to Start a Blog",
		URL:         "https://theartofsimple.net/startablog/",
		Image:       "https://unsplash.com/photos/IAytV8T2Qvc",
		BlogPost: BlogPosting{
			Backstory: "Every now and then, a reader will ask me a short-and-simple question that doesn’t exactly have a short-and-simple answer: 'How do I start a blog?'",
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
				MetaContext: context,
				MetaType:    "Blog",
				Name:        "How to Start a Blog",
				URL:         "https://theartofsimple.net/startablog/",
				Image:       "https://unsplash.com/photos/IAytV8T2Qvc",
				BlogPost: BlogPosting{
					Backstory: "Every now and then, a reader will ask me a short-and-simple question that doesn’t exactly have a short-and-simple answer: 'How do I start a blog?'",
				},
			},
			want: blog1.String(),
		},
		{
			name: "Test Case 1",
			fields: fields{
				MetaContext: context,
				MetaType:    "Blog",
				Name:        "Unit Testing In Go",
				URL:         "https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/",
				Image:       "https://testimage.com/test.jpeg",
				BlogPost: BlogPosting{
					Backstory: "To create basic tests for functions as well as basic tests for HTTP endpoints.",
				},
			},
			want: blog.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bg := Blog{
				MetaContext: tt.fields.MetaContext,
				MetaType:    tt.fields.MetaType,
				BlogPost:    tt.fields.BlogPost,
				Issn:        tt.fields.Issn,
				Name:        tt.fields.Name,
				URL:         tt.fields.URL,
				Image:       tt.fields.Image,
			}
			if got := bg.String(); got != tt.want {
				t.Errorf("Blog.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
