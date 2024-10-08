package orderedmap

import (
	"container/list"
	"iter"
	"sync"
)

// pair defines the type that's actually stored in the underlying linked list.
type pair[K comparable, V any] struct {
	key   K
	value V
}

// Map represents an ordered map.
//
// An ordered map preserves the inserting order when iterating.
//
// Underlying it's wrapping a sync.Map with a doubly linked list.
//
// The interface is intentionally kept almost the same with sync.Map to be used
// interchangeably.
//
// The zero value is an empty map ready to use.
// A map must not be copied after first use.
type Map[K comparable, V any] struct {
	lock sync.RWMutex

	l list.List
	m map[K]*list.Element
}

// Store stores the key value pair into the map.
func (m *Map[K, V]) Store(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	kv := &pair[K, V]{
		key:   key,
		value: value,
	}

	if element := m.m[key]; element != nil {
		// update existing value.
		element.Value = kv
		return
	}

	// insert new key-value pair to the back of the list.
	element := m.l.PushBack(kv)
	if m.m == nil {
		m.m = make(map[K]*list.Element)
	}
	m.m[key] = element
}

// Load loads key from the map.
//
// The ok result indicates whether the value is found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if element := m.m[key]; element != nil {
		return element.Value.(*pair[K, V]).value, true
	}
	return
}

// Delete deletes key from the map.
func (m *Map[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

// LoadAndDelete deletes the value for a key,
// returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if element := m.m[key]; element != nil {
		delete(m.m, key)
		m.l.Remove(element)
		return element.Value.(*pair[K, V]).value, true
	}
	return
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if element := m.m[key]; element != nil {
		kv := element.Value.(*pair[K, V])
		return kv.value, true
	}

	element := m.l.PushBack(&pair[K, V]{
		key:   key,
		value: value,
	})
	if m.m == nil {
		m.m = make(map[K]*list.Element)
	}
	m.m[key] = element
	return value, false
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// The order of the iteration preserves the original insertion order.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for e := m.l.Front(); e != nil; {
		kv := e.Value.(*pair[K, V])
		// Do it here instead of in for line to handle the special case of caller
		// deleting the key in f
		e = e.Next()

		m.lock.RUnlock()
		cont := f(kv.key, kv.value)
		m.lock.RLock()

		if !cont {
			break
		}
	}
}

// All returns iter.Seq2[key, value].
func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(key K, value V) bool) {
		m.Range(yield)
	}
}
