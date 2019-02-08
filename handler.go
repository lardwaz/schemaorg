package main

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	tplIndexDefault []byte
)

const typeArticle string = "Article"
const typeBlog string = "Blog"
const typeEvent string = "Event"
const typeMovie string = "Movie"
const typeJobPosting string = "JobPosting"
const typeRecipe string = "Recipe"
const typeWebPage string = "WebPage"
const typeWebSite string = "WebSite"

//Generator returbs a string
type Generator interface {
	String() (string, error)
}

//Templater interface forces other fucntion using it to implement template.Execute function
type Templater interface {
	Execute(wr io.Writer, data interface{}) error
}

//Handler returns an http.HandlerFunc
func Handler(tpl Templater, og OgParams, baseURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		webpageMeta := NewWebPageMeta(og.Title, og.URL, og.Description, og.Image, og.Keywords)
		schema, err := webpageMeta.String()

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("Handler(): String()")
			return
		}

		err = tpl.Execute(w, HTMLMeta{
			OG:           og,
			BaseURL:      baseURL,
			SchemaJSONld: schema,
		})

		if err != nil {
			log.WithFields(log.Fields{"error": err, "OG": og, "Schema": schema}).Warn("Handler(): tpl.Execute()")
			w.WriteHeader(http.StatusOK)
			w.Write(tplIndexDefault)
		}
	}
}

//DynamicHandler returns an http.HandlerFunc
func DynamicHandler(tpl Templater, baseURL string, og OgParams, g Generator) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		schema, err := g.String()

		if err != nil {
			return
		}

		err = tpl.Execute(w, HTMLMeta{
			OG:           og,
			BaseURL:      baseURL,
			SchemaJSONld: schema,
		})

		if err != nil {
			log.WithFields(log.Fields{"error": err, "OG": og, "Schema": schema}).Warn("Handler(): tpl.Execute()")
			w.WriteHeader(http.StatusOK)
			w.Write(tplIndexDefault)
		}
	}
}
