package web

import (
	"net/http"
	"strings"
	"sync"
)

// handler
type handler struct {
	mu           sync.Mutex
	routes       map[string]*route
	intercepts   []handleFunc
	filters      []handleFunc
	notFoundFunc handleFunc
}

func newHandler() *handler {
	handler := &handler{
		routes:       make(map[string]*route),
		notFoundFunc: notFound,
	}
	return handler
}

// match request url path
// If method is empty, support all method
// It is GET, only support GET method
func (this *handler) match(r *http.Request) (handleFunc, bool, bool) {
	method := r.Method
	path := muxPath(strings.ToLower(r.URL.Path))
	for pattern, route := range this.routes {
		if route.isFile && strings.HasPrefix(path, pattern) {
			return route.handle, true, true
		}
		if path == pattern && (route.method == "" || route.method == method) {
			return route.handle, true, false
		}
	}
	return nil, false, false
}

// serve http
func (this *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// match
	c := &Controller{r, w, make(map[string]interface{}), make(map[string]interface{})}
	handle, ok, isFile := this.match(r)
	if isFile {
		handle(c)
		return
	}
	if !ok {
		this.notFoundFunc(c)
		return
	}

	// intercept
	for _, h := range this.intercepts {
		if ok = h(c); ok {
			return
		}
	}

	// handle
	if handle(c) {
		return
	}

	// filter
	for _, h := range this.filters {
		if ok = h(c); ok {
			return
		}
	}
}

func (this *handler) route(pattern, method string, handle handleFunc) {
	pattern = muxPath(strings.ToLower(pattern))
	_, ok := this.routes[pattern]
	if ok {
		panic("HandleRoute Error: pattern is exist")
	}
	this.mu.Lock()
	this.routes[pattern] = &route{method: method, handle: handle}
	this.mu.Unlock()
}

func (this *handler) intercept(handle handleFunc) {
	this.mu.Lock()
	this.intercepts = append(this.intercepts, handle)
	this.mu.Unlock()
}

func (this *handler) filter(handle handleFunc) {
	this.mu.Lock()
	this.filters = append(this.filters, handle)
	this.mu.Unlock()
}

func (this *handler) static(pattern string, handle handleFunc) {
	pattern = muxPath(strings.ToLower(pattern))
	_, ok := this.routes[pattern]
	if ok {
		panic("HandleStatic Error: pattern is exist")
	}
	this.mu.Lock()
	this.routes[pattern] = &route{isFile: true, handle: handle}
	this.mu.Unlock()

}

func (this *handler) notFound(handle handleFunc) {
	this.mu.Lock()
	this.notFoundFunc = handle
	this.mu.Unlock()
}

// route
type route struct {
	method string
	isFile bool
	handle handleFunc
}

// return true is return
type handleFunc func(*Controller) bool
