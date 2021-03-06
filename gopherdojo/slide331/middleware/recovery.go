package middleware

import (
	"encoding/json"
	"github.com/aymanimam/hello-go/gopherdojo/slide331/errors"
	"log"
	"net/http"
	"runtime"
)

const panicText = "PANIC: %s\n%s"

// Recovery middleware
type Recovery struct{}

// ServeNext implements Middleware interface
func (r *Recovery) ServeNext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.NewOmikujiException(t, errors.OmikujiRecoveryErrorCode) // errors.New(t)
				case error:
					err = errors.NewOmikujiException(t.Error(), errors.OmikujiRecoveryErrorCode) // t
				default:
					err = errors.NewOmikujiException("unknown error", errors.OmikujiRecoveryUnknownErrorCode) // errors.New("unknown error")
				}

				// log the error
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				log.Printf(panicText, err, stack)

				// Return error response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				if jsonErr := json.NewEncoder(w).Encode(err); jsonErr != nil {
					log.Fatal(jsonErr)
				}
			}
		}()
		h.ServeHTTP(w, r)
	})
}

// NewRecovery Returns new recovery middleware
func NewRecovery() *Recovery {
	return &Recovery{}
}
