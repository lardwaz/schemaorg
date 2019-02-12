package schemaorg

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gocipe/schemaorg/schema"

	log "github.com/sirupsen/logrus"
)

var (
	tplIndexDefault []byte
)

//Templater interface forces other fucntion using it to implement template.Execute function
type Templater interface {
	Execute(wr io.Writer, data interface{}) error
}

//Handler returns an http.HandlerFunc
func Handler(tpl Templater, og OgParams, baseURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		webpageMeta := schema.NewWebPage(og.Title, og.Description, og.URL, og.Image, og.Keywords)
		schema := webpageMeta.String()

		err := tpl.Execute(w, HTMLMeta{
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
func DynamicHandler(tpl Templater, baseURL string, og OgParams, g fmt.Stringer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		schema := g.String()

		err := tpl.Execute(w, HTMLMeta{
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
