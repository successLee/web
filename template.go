package web

import (
	ht "html/template"
	"path/filepath"
	"strings"
	"sync"
)

type template struct {
	mu        sync.Mutex
	templates map[string]*ht.Template
	bases     []map[string]string
	funcMap   ht.FuncMap
}

func newTemplate() *template {
	t := &template{
		templates: make(map[string]*ht.Template),
		bases:     make([]map[string]string, 0),
		funcMap:   make(ht.FuncMap),
	}
	return t
}

func (this *template) base(prefix string, path string) error {
	path = filepath.Clean(path)
	this.mu.Lock()
	this.bases = append(this.bases, map[string]string{"prefix": prefix, "path": path})
	this.mu.Unlock()
	return nil
}

func (this *template) funcs(name string, f interface{}) {
	this.mu.Lock()
	this.funcMap[name] = f
	this.mu.Unlock()
}

func (this *template) load(path, suffix string) error {
	filePathList, err := getFilePathList(path, suffix)
	if err != nil {
		return err
	}

	for _, filePath := range filePathList {
		name := strings.Replace(filePath, filepath.Clean(path), "", 1)
		paths := make([]string, 0)
		isExist := false

		// skip
		for _, base := range this.bases {
			if base["path"] == filePath {
				isExist = true
				break
			}
		}

		if isExist {
			continue
		}

		// master
		for _, base := range this.bases {
			if strings.HasPrefix(name, base["prefix"]) {
				paths = append(paths, base["path"])
				break
			}
		}

		// content
		paths = append(paths, filePath)

		// template
		t, err := ht.New(filepath.Base(paths[0])).Funcs(this.funcMap).ParseFiles(paths...)
		if err != nil {
			return err
		}
		this.mu.Lock()
		this.templates[name] = t
		this.mu.Unlock()
	}
	return nil
}
