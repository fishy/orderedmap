package orderedmap

import (
	"sync"
)

var (
	_ mapInterface[any, any] = (*sync.Map)(nil)
	_ mapInterface[int, any] = (*Map[int, any])(nil)
)

// mapInterface defines the common interface between orderedmap.Map and sync.Map.
type mapInterface[K, V any] interface {
	Delete(key K)
	Load(key K) (value V, ok bool)
	LoadAndDelete(key K) (value V, loaded bool)
	LoadOrStore(key K, value V) (actual V, loaded bool)
	Range(f func(key K, value V) bool)
	Store(key K, value V)
}
