package gcpool

import (
	"log"
	"sync"
)

type Core struct {
	data map[string]interface{}
	lock *sync.RWMutex
	name string
}

func NewCore(name string) *Core {
	core := new(Core)
	core.name = name
	core.data = make(map[string]interface{})
	core.lock = new(sync.RWMutex)
	return core
}

func (this *Core) add(id string, obj interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	log.Printf("Core(%s) add(%s)", this.name, id)
	this.data[id] = obj
}

func (this *Core) len() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.data)
}

func (this *Core) get(id string, callback func(obj interface{}) error) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	// log.Printf("Core(%s) get(%s)", this.name, id)
	if obj, ok := this.data[id]; ok {
		callback(obj)
	}
}

func (this *Core) del(id string, callback func(obj interface{})) {
	this.lock.Lock()
	defer this.lock.Unlock()
	log.Printf("Core(%s) del(%s)", this.name, id)
	if obj, ok := this.data[id]; ok {
		callback(obj)
		delete(this.data, id)
	}
}
