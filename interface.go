package orderedmap

import (
	"sync"
)

var (
	_ Interface = (*sync.Map)(nil)
	_ Interface = (*Map)(nil)
)

// Interface defines the common interface between orderedmap.Map and sync.Map.
type Interface interface {
	Delete(key interface{})
	Load(key interface{}) (value interface{}, ok bool)
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Range(f func(key, value interface{}) bool)
	Store(key, value interface{})
}
