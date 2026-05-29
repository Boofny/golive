//Package middleware is the interface for middleware methods 
package middleware

import (
	"net/http"
	"strings"

	"github.com/Boofny/golive"
)


//CORS set to all origins * use only for testing not prod
func CORS() golive.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
      w.Header().Set("Access-Control-Allow-Credentials", "true")

  		if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return // stop here, don’t call next
      }

			next.ServeHTTP(w, r)
		})
	}
}

// CustomCORS allows as custom cors besides * origin
func CustomCORS(allowedOrigins... string) golive.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowed := false 

			for _, v := range allowedOrigins {
				if strings.EqualFold(v, origin){
					allowed = true
					break
				}
			}

			if allowed{
 				w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Vary", "Origin")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, X-CSRF-Token, Content-Type, Authorization")
        // w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

  		if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return // stop here, don’t call next
      }

			next.ServeHTTP(w, r)
		})
	}
}
