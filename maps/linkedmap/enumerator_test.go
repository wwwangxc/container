package linkedmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Any(t *testing.T) {
	m := testMap(5)

	got := m.Any(func(k, v string) bool {
		return k == "key_1" && v == "value_1"
	})
	assert.True(t, got)

	got = m.Any(func(k, v string) bool {
		return k == "key" && v == "value"
	})
	assert.False(t, got)
}

func BenchmarkMap_Any(b *testing.B) {
	m := testMap(1000)
	for i := 0; i < b.N; i++ {
		_ = m.Any(func(k, v string) bool {
			return false
		})
	}
}

func TestMap_All(t *testing.T) {
	m := testMap(5)

	got := m.All(func(k, v string) bool {
		return false
	})
	assert.False(t, got)

	got = m.All(func(k, v string) bool {
		return true
	})
	assert.True(t, got)
}

func BenchmarkMap_All(b *testing.B) {
	m := testMap(1000)

	for i := 0; i < b.N; i++ {
		_ = m.All(func(k, v string) bool {
			return true
		})
	}

	for i := 0; i < b.N; i++ {
		_ = m.All(func(k, v string) bool {
			return false
		})
	}
}

func TestMap_Find(t *testing.T) {
	m := testMap(5)

	k, v, exist := m.Find(func(k, v string) bool {
		return k == "key" && v == "value"
	})
	assert.Equal(t, "", k)
	assert.Equal(t, "", v)
	assert.False(t, exist)

	k, v, exist = m.Find(func(k, v string) bool {
		return k == "key_1" && v == "value_1"
	})
	assert.Equal(t, "key_1", k)
	assert.Equal(t, "value_1", v)
	assert.True(t, exist)
}

func BenchmarkMap_Find(b *testing.B) {
	m := testMap(1000)

	for i := 0; i < b.N; i++ {
		_, _, _ = m.Find(func(k, v string) bool {
			return k == "key" && v == "value"
		})
	}

	for i := 0; i < b.N; i++ {
		_, _, _ = m.Find(func(k, v string) bool {
			return k == "key_1" && v == "value_1"
		})
	}
}

func TestMap_Select(t *testing.T) {
	m := testMap(5)

	values := m.Select(func(k, v string) bool {
		return k == "key" && v == "value"
	})
	assert.Equal(t, []string(nil), values)

	values = m.Select(func(k, v string) bool {
		return true
	})
	assert.Equal(t, []string{"value_0", "value_1", "value_2", "value_3", "value_4"}, values)
}

func BenchmarkMap_Select(b *testing.B) {
	m := testMap(1000)

	for i := 0; i < b.N; i++ {
		_ = m.Select(func(k, v string) bool {
			return true
		})
	}
}

func testMap(length int) *Map[string, string] {
	m := New[string, string]()
	for i := 0; i < length; i++ {
		m.Put(fmt.Sprintf("key_%d", i), fmt.Sprintf("value_%d", i))
	}

	return m
}
