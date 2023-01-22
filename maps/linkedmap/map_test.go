package linkedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Put(t *testing.T) {
	m := New[string, string]()
	assert.Equal(t, 0, len(m.m))
	assert.Equal(t, 0, m.list.Size())

	m.Put("key_1", "value_1")
	assert.Equal(t, 1, len(m.m))
	assert.Equal(t, 1, m.list.Size())
	assert.True(t, m.Contains("key_1"))
}

func TestMap_Get(t *testing.T) {
	m := New[string, string]()

	v, exist := m.Get("key")
	assert.Equal(t, "", v)
	assert.False(t, exist)

	m.Put("key", "value")
	v, exist = m.Get("key")
	assert.Equal(t, "value", v)
	assert.True(t, exist)
}

func TestMap_Delete(t *testing.T) {
	m := New[string, string]()
	m.Put("key_1", "value_1")
	assert.Equal(t, 1, len(m.m))
	assert.Equal(t, 1, m.list.Size())
	assert.True(t, m.Contains("key_1"))

	m.Delete("key")
	assert.Equal(t, 1, len(m.m))
	assert.Equal(t, 1, m.list.Size())
	assert.True(t, m.Contains("key_1"))

	m.Delete("key_1")
	assert.Equal(t, 0, len(m.m))
	assert.Equal(t, 0, m.list.Size())
	assert.False(t, m.Contains("key_1"))
}

func TestMap_Contains(t *testing.T) {
	m := New[string, string]()
	m.Put("key_1", "value_1")
	assert.False(t, m.Contains("key"))
	assert.True(t, m.Contains("key_1"))
}

func TestMap_Keys(t *testing.T) {
	m := New[string, string]()
	keys := m.Keys()
	assert.Equal(t, 0, len(keys))

	m.Put("key_1", "value_1")
	m.Put("key_2", "value_2")
	m.Put("key_3", "value_3")
	keys = m.Keys()
	assert.Equal(t, 3, len(keys))
}
