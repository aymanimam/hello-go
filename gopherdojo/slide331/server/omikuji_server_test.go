package server

import (
	"github.com/aymanimam/hello-go/gopherdojo/slide331/errors"
	"github.com/aymanimam/hello-go/gopherdojo/slide331/middleware"
	"github.com/aymanimam/hello-go/gopherdojo/slide331/omikuji"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

var validResponseRegex = regexp.MustCompile(`^\\{"omikuji":"(大吉|吉|中吉|小吉|半吉|末吉|末小吉|凶|小凶|半凶|末凶|大凶)"\\}$`)

// Mock omikuji.Dispatcher
type MockPanicOmikujiDispatcher struct{}

func (omikujis *MockPanicOmikujiDispatcher) GetNextOmikuji() omikuji.Omikuji {
	msg := "MockPanicOmikujiDispatcher Error!"
	errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
	return omikuji.Omikuji{}
}

// ---------------

func TestOmikujiHandler(t *testing.T) {
	// Initialize
	initialize()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/omikuji", nil)
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
	omikujiDispatcher = &MockPanicOmikujiDispatcher{}

	rec := middleware.NewRecovery()
	mw := middleware.With(http.HandlerFunc(omikujiHandler), rec)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/omikuji", nil)

	mw.ServeHTTP(w, r)

	rw := w.Result()

	defer rw.Body.Close()

	if rw.StatusCode != http.StatusInternalServerError {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	errStr := string(b)
	if !strings.Contains(errStr, "[MockPanicOmikujiDispatcher Error!") {
		t.Fatalf("unexpected response: %s", errStr)
	}
}
