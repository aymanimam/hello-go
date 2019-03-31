package server

import (
	"github.com/aymanimam/hello-go/gopherdojo/slide331/omikuji"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var validResponseRegex = regexp.MustCompile(`^\\{"omikuji":"(大吉|吉|吉|中吉|小吉|半吉|末吉|末小吉|凶|小凶|半凶|末凶|大凶)"\\}$`)

func TestOmikujiHandler(t *testing.T) {
	// Initialize
	initialize()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	omikujiHandler(w, r)
	rw := w.Result()

	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	if validResponseRegex.Match(b) {
		t.Fatalf("unexpected response: %s", string(b))
	}
}

func TestOmikujiHandlerErrorResponse(t *testing.T) {
	// Initialize
	periodChecker = nil
	randomizer = omikuji.GetOmikujiRandomizer()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	omikujiHandler(w, r)
	rw := w.Result()

	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	if validResponseRegex.Match(b) {
		t.Fatalf("unexpected response: %s", string(b))
	}
}
