package ordered_sync_map

import (
	"container/list"
	"sync"
)

type mapElement struct {
	key   interface{}
	value interface{}
}

// Map is a thread safe and ordered implementation of standard map.
type Map struct {
	mp  map[interface{}]*list.Element
	mu  sync.RWMutex
	dll *list.List
}

// New returns an initialized Map.
func New() *Map {
	m := new(Map)
	m.mp = make(map[interface{}]*list.Element)
	m.dll = list.New()
	return m
}

// Get returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map) Get(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	v, ok := m.mp[key]
	if !ok {
		m.mu.RUnlock()
		return nil, false
	}
	m.mu.RUnlock()
	me := v.Value.(mapElement)
	return me.value, ok
}

// Put sets the value for the given key.
// It will replace the value if the key already exists in the map
// even if the values are same.
func (m *Map) Put(key interface{}, val interface{}) {
	m.mu.Lock()
	if e, ok := m.mp[key]; !ok {
		m.mp[key] = m.dll.PushFront(mapElement{key: key, value: val})
	} else {
		e.Value = mapElement{key: key, value: val}
	}
	m.mu.Unlock()
}

// Delete deletes the value for a key.
// It returns a boolean indicating weather the key existed and it was deleted.
func (m *Map) Delete(key interface{}) bool {
	m.mu.Lock()
	e, ok := m.mp[key]
	if !ok {
		m.mu.Unlock()
		return false
	}

	m.dll.Remove(e)
	delete(m.mp, key)
	m.mu.Unlock()
	return true
}

// UnorderedRange will range over the map in an unordered sequence.
// This is same as ranging over a map using the "for range" syntax.
func (m *Map) UnorderedRange(f func(key interface{}, value interface{})) {
	m.mu.RLock()
	for k, v := range m.mp {
		f(k, v.Value.(mapElement).value)
	}
	m.mu.RUnlock()
}

// OrderedRange will range over the map in ab ordered sequence.
// This will probably be slower than UnorderedRange()
func (m *Map) OrderedRange(f func(key interface{}, value interface{})) {
	m.mu.RLock()
	if m.dll.Len() == 0 {
		m.mu.RUnlock()
		return
	}

	cur := m.dll.Back()
	for cur != nil {
		me := cur.Value.(mapElement)
		f(me.key, me.value)
		cur = cur.Prev()
	}

	m.mu.RUnlock()
}
