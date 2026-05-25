package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
type wrappedWrite struct{
	http.ResponseWriter
	satusCode int
	wroteHeader bool
}

func (w *wrappedWrite)WriteHeader(statusCode int){ //need to look into this 
	if w.wroteHeader {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.satusCode = statusCode
	w.wroteHeader = true
}//is this recursive??????

func (w *wrappedWrite) Write(b []byte) (int, error) {
  if !w.wroteHeader {
    w.WriteHeader(http.StatusOK)
  }
  return w.ResponseWriter.Write(b)
}

//NOTE: commenting this out in order to test something if want to go back to optional logger uncommnt this and see the golive file

// func Logger() Middleware{
// 	return func (next http.Handler)http.Handler  { //og name is Logging() new name Logger()
// 		redH:= "\033[31m"
// 		greenH := "\033[32m"
// 		reset := "\033[0m"
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 			start := time.Now()
//
// 			wrapped := &wrappedWrite{
// 				ResponseWriter: w,
// 				satusCode: http.StatusOK,
// 			}
// 			next.ServeHTTP(wrapped, r)
// 			code := wrapped.satusCode
// 			if code >= 400 && code <= 599{ 
// 				fmt.Print("\033[31m >>> \033[0m") //error
// 				log.Println(redH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start))
// 			}else{
// 				fmt.Print("\033[32m >>> \033[0m") //good
// 				log.Println(greenH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start)) 
// 			}
// 		})
// 	}
// }

// func Logger(next http.Handler)http.Handler  { //og name is Logging() new name Logger()
// 	redH:= "\033[31m"
// 	greenH := "\033[32m"
// 	reset := "\033[0m"
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 		start := time.Now()
//
// 		wrapped := &wrappedWrite{
// 			ResponseWriter: w,
// 			satusCode: http.StatusOK,
// 		}
// 		next.ServeHTTP(wrapped, r)
// 		code := wrapped.satusCode
// 		if code >= 400 && code <= 599{ 
// 			fmt.Print("\033[31m >>> \033[0m") //error
// 			log.Println(redH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start))
// 		}else{
// 			fmt.Print("\033[32m >>> \033[0m") //good
// 			log.Println(greenH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start)) 
// 		}
// 	})
// }

func Logger() Middleware{
	return func (next http.Handler) http.Handler {
		// Color codes
		red := "\033[31m"
		green := "\033[32m"
		yellow := "\033[33m"
		magenta := "\033[35m"
		reset := "\033[0m"

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrapped := &wrappedWrite{
				ResponseWriter: w,
				satusCode:      http.StatusOK,
			}

			next.ServeHTTP(wrapped, r)

			code := wrapped.satusCode
			duration := time.Since(start)

			// Pick color based on status code range
			var color string
			switch {
			case code >= 200 && code < 300:
				color = green   // Success
			case code >= 300 && code < 400:
				color = yellow  // Redirect
			case code >= 400 && code < 500:
				color = red     // Client error
			case code >= 500:
				color = magenta // Server error
			default:
				color = reset   // Fallback
			}

			// Output format (your original style)
			// fmt.Print("[ OUTPUT ]:")
			// fmt.Printf("%s >>> %s", color, reset)
			// log.Println(color, wrapped.satusCode, reset, r.Method, "[",r.URL.Path,"]", "|",duration)
			log.Printf("[ OUTPUT ]: >>> %s%d%s  %s [ %s ] | %s",
					color, wrapped.satusCode, reset,
					r.Method,
					r.URL.Path,
					duration,
			)
		})
	}
}

// UniversalLogger NOTE: used for other projects
func UniversalLogger(next http.Handler) http.Handler {
	// Color codes
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	magenta := "\033[35m"
	reset := "\033[0m"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWrite{
			ResponseWriter: w,
			satusCode:      http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		code := wrapped.satusCode
		duration := time.Since(start)

		// Pick color based on status code range
		var color string
		switch {
		case code >= 200 && code < 300:
			color = green   // Success
		case code >= 300 && code < 400:
			color = yellow  // Redirect
		case code >= 400 && code < 500:
			color = red     // Client error
		case code >= 500:
			color = magenta // Server error
		default:
			color = reset   // Fallback
		}

		// Output format (your original style)
		fmt.Print("[ OUTPUT ]:")
		fmt.Printf("%s >>> %s", color, reset)
		log.Println(color, wrapped.satusCode, reset, r.Method, "[",r.URL.Path,"]", "|",duration)
	})
}
