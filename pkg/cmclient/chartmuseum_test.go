package cmclient

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

type values map[string]string

// setup sets up a test HTTP server along with a ChartMuseum Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()
	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the ChartMuseum client being tested and is
	// configured to use test server.
	client, _ = NewClient(server.URL+"/", nil, "user", "pass")
	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}
func testBody(t *testing.T, r *http.Request, want []byte) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if !reflect.DeepEqual(b, want) {
		t.Errorf("request Body got %v want %v", len(b), len(want))
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}
