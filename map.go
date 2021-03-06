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
	kv := &pair{key, value}
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
	kv := element.Value.(*pair)
	return kv.Value, true
}

// Delete deletes key from the map.
//
// key must be hashable.
func (m *Map) Delete(key interface{}) {
	value, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return
	}
	m.l.Remove(value.(*list.Element))
}

// LoadAndDelete deletes the value for a key,
// returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	value, loaded = m.m.LoadAndDelete(key)
	if !loaded {
		return
	}
	element := value.(*list.Element)
	m.l.Remove(element)
	return element.Value.(*pair).Value, true
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
//
// key must be hashable.
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	element := m.getElement(key)
	if element != nil {
		kv := element.Value.(*pair)
		return kv.Value, true
	}
	kv := &pair{key, value}
	element = m.l.PushBack(kv)
	m.m.Store(key, element)
	return value, false
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// The order of the iteration preserves the original insertion order.
func (m *Map) Range(f func(key, value interface{}) bool) {
	for e := m.l.Front(); e != nil; {
		kv := e.Value.(*pair)
		// Do it here instead of in for line to handle the special case of caller
		// deleting the key in f
		e = e.Next()
		if !f(kv.Key, kv.Value) {
			break
		}
	}
}
