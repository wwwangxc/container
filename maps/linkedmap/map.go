package linkedmap

import "github.com/wwwangxc/container"

var _ (container.Map[string, string]) = (*Map[string, string])(nil)

// Put the key and the value
func (m *Map[K, V]) Put(k K, v V) {
	if m.Contains(k) {
		m.list.Remove(m.list.IndexOf(k))
	}

	m.m[k] = v
	m.list.Append(k)
}

// Get value by key
func (m *Map[K, V]) Get(k K) (V, bool) {
	var v V

	if !m.Contains(k) {
		return v, false
	}

	return m.m[k], true
}

// Delete value by ke
func (m *Map[K, V]) Delete(k K) {
	if !m.Contains(k) {
		return
	}

	delete(m.m, k)
	m.list.Remove(m.list.IndexOf(k))
}

// Contains all given keys
func (m *Map[K, V]) Contains(keys ...K) bool {
	if m.IsEmpty() {
		return false
	}

	for _, k := range keys {
		if _, ok := m.m[k]; !ok {
			return false
		}
	}

	return true
}

// Keys return all keys in the map
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, m.Size())
	m.list.Each(func(_ int, k K) bool {
		keys = append(keys, k)
		return true
	})

	return keys
}
