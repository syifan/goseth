package seth

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

var _once sync.Once
var _instance *registryImpl

// TypeRegistry keeps track of all the types available in the current
// application.
//
// The TypeRegistry implementation is following this answer on stackoverflow.
// https://stackoverflow.com/a/34722791/1709930
type TypeRegistry interface {
	Register(interface{})
	GetType(name string) reflect.Type
}

// GetTypeRegistry get the singleton instance of the TypeRegistry
func GetTypeRegistry() TypeRegistry {
	_once.Do(func() {
		_instance = &registryImpl{
			dict: make(map[string]reflect.Type),
		}
	})

	return _instance
}

type registryImpl struct {
	dict map[string]reflect.Type
}

func (r *registryImpl) Register(typedNil interface{}) {
	t := reflect.TypeOf(typedNil).Elem()
	name := t.PkgPath() + "." + t.Name()
	r.dict[name] = t
}

func (r *registryImpl) GetType(name string) reflect.Type {
	return r.dict[name]
}

func BuildRegistry() {
	i := 0
	found := true
	for found {
		pc, file, line, ok := runtime.Caller(i)
		fmt.Println(pc, file, line, ok)
		found = ok
		i++
	}
}
