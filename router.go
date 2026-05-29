package golive

import (
	"net/http"
	"strings"
)

type route struct{
	method string
	pattern string 
	parts []string
	handler HandleFunc 
}

type Router struct{
	routes []route
	prefixRoutes []prefixroute
}

type prefixroute struct{
	prefix string
	handler http.Handler
}

// matchPath checks if a request path matches a pattern and extracts params.
// Pattern segments wrapped in {} are wildcards: /user/{id}/posts/{postId}
func matchPath(parts []string, pathSegments []string) (map[string]string, bool) {
	if len(parts) != len(pathSegments) {
		return nil, false
	}
	params := make(map[string]string)
	for i, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			key := part[1 : len(part)-1]
			params[key] = pathSegments[i]
		} else if part != pathSegments[i] {
			return nil, false
		}
	}
	return params, true
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Trim leading slash and split into segments
	pathSegments := strings.Split(strings.Trim(req.URL.Path, "/"), "/")

	for _, route := range r.routes {
		if route.method != req.Method {
			continue
		}
		params, ok := matchPath(route.parts, pathSegments)
		if !ok {
			continue
		}
		ctx := &Context{Writer: w, Request: req, params: params}
		if err := route.handler(ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

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
