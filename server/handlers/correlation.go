package handlers

import "net/http"
import "github.com/google/uuid"

type correlationHandler struct {
	next http.Handler
}

func (c *correlationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Request-ID") == "" {
		r.Header.Set("X-Request-ID", uuid.New().String())
	}

	c.next.ServeHTTP(rw, r)
}

// NewCorrelationHandler is a middleware handler which appends a X-Request-ID
// header if one is not already set before calling the next handler in the chain
func NewCorrelationHandler(next http.Handler) http.Handler {
	return &correlationHandler{next: next}
}
