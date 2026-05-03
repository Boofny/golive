package goLive

import "net/http"

type route struct{
	method string
	pattern string 
	handler FunctionHandler
}

type Router struct{
	routes []route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.method == req.Method && route.pattern == req.URL.Path {
			ctx := &Context{Writer: w, Request: req}
			if err := route.handler(ctx); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.NotFound(w, req)
}
