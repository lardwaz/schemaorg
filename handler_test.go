package schemaorg_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"github.com/gorilla/mux"
	schemaorg "go.lsl.digital/lardwaz/schemaorg"
	"go.lsl.digital/lardwaz/schemaorg/schema"
	"github.com/stretchr/testify/assert"
)

var MockOgraph = schemaorg.NewOgParams("This is title", "http://localhost:8888/a", "This is a description", "https://via.placeholder.com/450/0000FF/808080%20?Text=Digital.com%20C/O%20https://placeholder.com/", "test, food")
var MockBaseURL = "http://localhost:8888/a"

var MockOgraphMovie = MockMovieOg()

func MockMovieOg() schema.Movie {
	directors := []string{"Steven Spielbery", "James Cameron", "Ron Howard"}
	actors := []string{"Christian Bale", "Johny Depp"}
	return schema.NewMovie("Movie1", "Description is written here", "", "imageurl.com/img.jpeg", directors, actors)
}

func MockEndPointDynamic() http.HandlerFunc {
	tmpl := template.New("Test")
	tmpl, _ = tmpl.Parse("OG: {{.OG}} , baseurl: {{.BaseURL}}, schema: {{.SchemaJSONld}}")

	og := schemaorg.NewOgParams(MockOgraphMovie.Name, MockOgraphMovie.URL, MockOgraphMovie.Description, MockOgraphMovie.Image, "")

	return schemaorg.DynamicHandler(tmpl, MockBaseURL, og, MockOgraphMovie)
}

func MockEndPointFunc() http.HandlerFunc {
	tmpl := template.New("Test")
	tmpl, _ = tmpl.Parse("OG: {{.OG}} , baseurl: {{.BaseURL}}")

	return schemaorg.Handler(tmpl, MockOgraph, MockBaseURL)
}

func MockRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/handlertest", MockEndPointFunc()).Methods("GET")
	router.HandleFunc("/dynamichandlertest", MockEndPointDynamic()).Methods("GET")
	return router
}

func TestHandler(t *testing.T) {

	request, _ := http.NewRequest("GET", "/handlertest", nil)
	response := httptest.NewRecorder()
	MockRouter().ServeHTTP(response, request)

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	boolState := strings.Contains(bodyString, MockOgraph.Title)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockOgraph.URL)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockOgraph.Description)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockOgraph.Image)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockOgraph.Keywords)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockBaseURL)
	assert.True(t, boolState)

	assert.Equal(t, 200, response.Code)
}

func TestDynamicHandler(t *testing.T) {
	tmpl := template.New("Test")
	tmpl, _ = tmpl.Parse("OG: {{.OG}} , baseurl: {{.BaseURL}}, Schema: {{.schema}}")

	request, _ := http.NewRequest("GET", "/dynamichandlertest", nil)
	response := httptest.NewRecorder()
	MockRouter().ServeHTTP(response, request)

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	boolState := strings.Contains(bodyString, MockOgraphMovie.Name)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockOgraphMovie.URL)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockOgraphMovie.Description)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockOgraphMovie.Image)
	assert.True(t, boolState)
	boolState = strings.Contains(bodyString, MockBaseURL)
	assert.True(t, boolState)

	boolState = strings.Contains(bodyString, MockOgraphMovie.String())
	assert.True(t, boolState)

	assert.Equal(t, 200, response.Code)
}
