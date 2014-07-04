package web

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Request  *http.Request
	Response http.ResponseWriter
	Result   map[string]interface{}
	Internal map[string]interface{}
}

func notFound(c *Controller) bool {
	http.NotFound(c.Response, c.Request)
	return false
}

func (c *Controller) NotFound() bool {
	return notFound(c)
}

func (c *Controller) Redirect(url string) bool {
	http.Redirect(c.Response, c.Request, url, http.StatusFound)
	return false
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
	}
	http.SetCookie(c.Response, cookie)
}
