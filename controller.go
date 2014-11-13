package web

import (
	"encoding/json"
	"net/http"
	"bytes"
)

type Controller struct {
	Request  *http.Request
	Response http.ResponseWriter
	Result   map[string]interface{}
	Internal map[string]interface{}
}

func notFound(c *Controller) bool {
	http.NotFound(c.Response, c.Request)
	return true
}

func methodNotAllowed(c *Controller) bool {
	http.Error(c.Response, http.ErrBodyNotAllowed.Error(), http.StatusMethodNotAllowed)
	return true
}

func (c *Controller) NotFound() bool {
	return notFound(c)
}

func (c *Controller) MethodNotAllowed() bool {
	return methodNotAllowed(c)
}

func (c *Controller) Redirect(url string) bool {
	http.Redirect(c.Response, c.Request, url, http.StatusFound)
	return true
}

func (c *Controller) Render(name string) bool {
	if name[0] != '/' {
		name = "/" + name
	}
	t, ok := defaultTemplate.templates[name]
	if !ok {
		return c.NotFound()
	}
	t.Execute(c.Response, c.Result)
	return false
}

func (c *Controller) RenderJson() bool {
	b, _ := json.Marshal(c.Result)
	c.Response.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Response.Write(b)
	return false
}

func (c *Controller) RenderJsonP(callback string) bool {
	buffer := bytes.NewBufferString("")
	buffer.WriteString(callback)
	buffer.WriteString("(")
	b, _ := json.Marshal(c.Result)
	buffer.Write(b)
	buffer.WriteString(")")
	c.Response.Write(buffer.Bytes())
	return false
}

// cookie
func (c *Controller) GetCookie(name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

func (c *Controller) SetCookie(name, value string) {
	cookie := &http.Cookie{
		Name: name, Value: value,
		Path:       defaultCookieConfig.Path,
		Domain:     defaultCookieConfig.Domain,
		Expires:    defaultCookieConfig.Expires,
		RawExpires: defaultCookieConfig.RawExpires,
		MaxAge:     defaultCookieConfig.MaxAge,
		Secure:     defaultCookieConfig.Secure,
		HttpOnly:   defaultCookieConfig.HttpOnly,
		Raw:        defaultCookieConfig.Raw,
		Unparsed:   defaultCookieConfig.Unparsed,
	}
	http.SetCookie(c.Response, cookie)
}

func (c *Controller) DelCookie(name string) {
	cookie := &http.Cookie{
		Name:   name,
		MaxAge: -1,
		Path:   defaultCookieConfig.Path,
		Domain: defaultCookieConfig.Domain,
	}
	http.SetCookie(c.Response, cookie)
}
