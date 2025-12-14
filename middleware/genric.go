// TODO: next functions to work on is EnsureAdmin and Authorization for requests
package middleware

import (
	// "log"
	"net/http"
	"strings"
	// "strings"
)

// func EnsureAdmin(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Checking if user is admin")
// 		if !strings.Contains(r.Header.Get("Authorization"), "Admin") {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
//
// func LoadUser(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// log.Println("Loading user")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func CORS2 () Middleware{
// 	return func (next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Header().Set("Access-Control-Allow-Origin", "*")//will need this bro bro
// 			log.Println("Enabling CORS")
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// type Middleware func(http.Handler) http.Handler //just to remind what middleware defines

//CORS set to all origins * use only for testing
func CORS() Middleware {
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

//TODO: need to add an option for multi origin
//should prob be an array aka slice that contains multiple origins 
//can remake this function but with this in mind

	// e.Use(middleware.Recover())
	//
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins:     []string{"https://*", "http://*"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	// 	AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	AllowCredentials: true,
	// 	MaxAge:           300,
	// }))

// CustomCORS allows as custom cors besides * origin
func CustomCORS(allowedOrigins... string) Middleware {
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
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
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

// func CheckPermissions(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Checking Permissions")
// 		next.ServeHTTP(w, r)
// 	})
// }
