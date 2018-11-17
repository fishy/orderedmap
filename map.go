package orderedmap

import (
	"container/list"
	"sync"
)

// pair defines the type that's actually stored in the underlying linked list.
type pair struct {
	Key   interface{}
	Value interface{}
}

// Map represents an ordered map.
//
// An ordered map preserves the inserting order when iterating.
//
// Underlying it's wrapping a sync.Map with a doubly linked list.
//
// The interface is intentionally kept the same with sync.Map to be used
// interchangeably.
//
// The zero value is an empty map ready to use.
// A map must not be copied after first use.
type Map struct {
	l list.List
	m sync.Map
}

func (m *Map) getElement(key interface{}) *list.Element {
	element, ok := m.m.Load(key)
	if !ok {
		return nil
	}
	return element.(*list.Element)
}

// Store stores the key value pair into the map.
//
// key must be hashable.
func (m *Map) Store(key, value interface{}) {
	kv := pair{key, value}
	element := m.getElement(key)
	if element != nil {
		// update existing value.
		element.Value = kv
		return
	}
	// insert new key-value pair to the back of the list.
	element = m.l.PushBack(kv)
	m.m.Store(key, element)
}

// Load loads key from the map.
//
// key must be hashable.
//
// The ok result indicates whether the value is found in the map.
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	element := m.getElement(key)
	if element == nil {
		return nil, false
	}
	kv := element.Value.(pair)
	return kv.Value, true
}

// Delete deletes key from the map.
//
// key must be hashable.
func (m *Map) Delete(key interface{}) {
	element := m.getElement(key)
	if element == nil {
		return
	}
	m.l.Remove(element)
	m.m.Delete(key)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
//
// key must be hashable.
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	element := m.getElement(key)
	if element != nil {
		kv := element.Value.(pair)
		return kv.Value, true
	}
	kv := pair{key, value}
	element = m.l.PushBack(kv)
	m.m.Store(key, element)
	return value, false
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// The order of the iteration preserves the original insertion order.
func (m *Map) Range(f func(key, value interface{}) bool) {
	for e := m.l.Front(); e != nil; e = e.Next() {
		kv := e.Value.(pair)
		if !f(kv.Key, kv.Value) {
			break
		}
	}
}
