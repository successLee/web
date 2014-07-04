package web

import (
	"errors"
	"sync"
)

type Idb interface {
	Get(name, id string) (interface{}, error)
	Put(name, id string, model interface{}) error
}

// memory
type memory struct {
	mu    sync.Mutex
	store map[string]map[string]interface{}
}

func newMemoryDB() Idb {
	var db Idb = &memory{store: make(map[string]map[string]interface{})}
	return db
}

func (this *memory) Get(name string, id string) (interface{}, error) {
	m, ok := this.store[name]
	if ok {
		model, ok := m[id]
		if ok {
			return model, nil
		}
	}
	return nil, errors.New("not found")
}

func (this *memory) Put(name, id string, model interface{}) error {
	m, ok := this.store[name]
	this.mu.Lock()
	if !ok {
		this.store[name] = map[string]interface{}{id: model}
	} else {
		m[id] = model
	}
	this.mu.Unlock()
	return nil
}
