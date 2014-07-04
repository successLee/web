package web

import (
	"fmt"
	"sync"
)

type local struct {
	mu          sync.Mutex
	defaultLang string
	locals      map[string]map[string]string
}

func newLocal() *local {
	return &local{defaultLang: "en_US", locals: make(map[string]map[string]string)}
}

func (this *local) setDefault(lang string) {
	this.mu.Lock()
	this.defaultLang = lang
	this.mu.Unlock()
}

func (this *local) setMessage(lang, key, value string) {
	_, ok := this.locals[lang]
	this.mu.Lock()
	if !ok {
		this.locals[lang] = make(map[string]string)

	}
	_, ok = this.locals[lang][key]
	if !ok {
		this.locals[lang][key] = value
	}
	this.mu.Unlock()
}

func (this *local) value(lang, key string, values ...interface{}) string {
	defaultLang := this.defaultLang

	// get default lang
	_, ok := this.locals[lang]
	if ok {
		defaultLang = lang
	} else {
		subDefaultLang, ok := this.locals[lang[:2]]
		if ok {
			subLang, ok := subDefaultLang["default"]
			if ok {
				defaultLang = subLang
			}
		}
	}

	// get message
	locals, ok := this.locals[defaultLang]
	if !ok {
		return key
	}

	value, ok := locals[key]
	if !ok {
		return key
	}

	if len(values) <= 0 {
		return value
	}

	value = fmt.Sprintf(value, values...)
	return value
}
