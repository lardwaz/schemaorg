package schemaorg

import (
	"fmt"
	"net/http"
)

type Resolver func(id string) fmt.Stringer

func OgHandler(tpl string, data fmt.Stringer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//writeheader text/html
		w.WriteHeader(http.StatusOK)
		w.Write(tplIndexDefault)
		panic("Not implemented")
	}
}

func OgSchemaHandler(tpl string, resolve Resolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
