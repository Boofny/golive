package goLive

import (
	"net/http"
	"strings"
)

type route struct{
	method string
	pattern string 
	handler FunctionHandler
}

type Router struct{
	routes []route
	prefixRoutes []prefixroute
}

type prefixroute struct{
	prefix string
	handler http.Handler
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
	// prefix match for static dirs
	for _, pr := range r.prefixRoutes {
		if strings.HasPrefix(req.URL.Path, pr.prefix) {
			pr.handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (g *GoLive)ServeDir(urlPath, dirPath string) error {
	fs := http.FileServer(http.Dir(dirPath))
	g.router.prefixRoutes = append(g.router.prefixRoutes, prefixroute{
		prefix: urlPath,
		handler: http.StripPrefix(urlPath, fs),
	})
	return nil
}
