package middleware

import (
	"net/http"

	goliveMiddleware "github.com/Boofny/golive"
	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

// RateLimit dead simple middleware using golang's rate limiting package : for now its a global limiter
func RateLimit(Request rate.Limit, TokenPerRequest int) goliveMiddleware.Middleware {
	limiter = rate.NewLimiter(Request, TokenPerRequest)
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "rate limit reached", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
