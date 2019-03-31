package middleware

import (
	"errors"
	"log"
	"net/http"
	"runtime"
)

const panicText = "PANIC: %s\n%s"

// Recovery middleware
type Recovery struct{}

func (r *Recovery) ServeNext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}
				// log the error
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				log.Printf(panicText, err, stack)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func NewRecovery() *Recovery {
	return &Recovery{}
}
