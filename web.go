package web

import (
	"net/http"
)

// private
var (
	defaultHandler      = newHandler()
	defaultCookieConfig = newCookieConfig()
	defaultTemplate     = newTemplate()
	defaultLocal        = newLocal()
)

// public
var (
	Handle   = newWebHandle()
	Cookie   = newWebCookie()
	Template = newWebTemplate()
	DB       = newWebDB()
	Local    = newWebLocal()
)

// web route
type webHandle struct{}

func newWebHandle() *webHandle {
	return &webHandle{}
}

func (this *webHandle) Route(pattern, method string, handle handleFunc) {
	defaultHandler.route(pattern, method, handle)
}

func (this *webHandle) Intercept(handle handleFunc) {
	defaultHandler.intercept(handle)
}

func (this *webHandle) Filter(handle handleFunc) {
	defaultHandler.filter(handle)
}

func (this *webHandle) Static(pattern string, handle handleFunc) {
	defaultHandler.static(pattern, handle)
}

func (this *webHandle) NotFound(handle handleFunc) {
	defaultHandler.notFound(handle)
}

func (this *webHandle) MethodNotAllowed(handle handleFunc) {
	defaultHandler.methodNotAllowed(handle)
}

// web cookie
type webCookie struct{}

func newWebCookie() *webCookie {
	return &webCookie{}
}

func (this *webCookie) Reset(cookieConfig *CookieConfig) {
	cc := cookieConfig
	defaultCookieConfig = cc
}

// web template
type webTemplate struct{}

func newWebTemplate() *webTemplate {
	return &webTemplate{}
}

func (this *webTemplate) Reset() {
	t := newTemplate()
	defaultTemplate = t
}

func (this *webTemplate) Base(prefix string, path string) error {
	return defaultTemplate.base(prefix, path)
}

func (this *webTemplate) Funcs(name string, f interface{}) {
	defaultTemplate.funcs(name, f)
}

func (this *webTemplate) Load(path, suffix string) error {
	return defaultTemplate.load(path, suffix)
}

// web db
type webDB struct{}

func newWebDB() *webDB {
	return &webDB{}
}

func (this *webDB) newMemory() Idb {
	return newMemoryDB()
}

// web local
type webLocal struct{}

func newWebLocal() *webLocal {
	return &webLocal{}
}

func (this *webLocal) Reset() {
	local := newLocal()
	defaultLocal = local
}

func (this *webLocal) setDefault(lang string) {
	defaultLocal.setDefault(lang)
}

func (this *webLocal) setMessage(lang, key, value string) {
	defaultLocal.setMessage(lang, key, value)
}

func (this *webLocal) value(lang, key string, values ...interface{}) string {
	return defaultLocal.value(lang, key, values...)
}

// run server in addr
func Run(addr string) error {
	return http.ListenAndServe(addr, defaultHandler)
}
