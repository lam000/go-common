package paladin

import (
	"strings"
	"sync/atomic"
)

type Map struct {
	values atomic.Value
}

func (m *Map) Store(values map[string]*Value) {
	dst := make(map[string]*Value, len(values))
	for k, v := range values {
		dst[keyNamed(k)] = v
	}
	m.values.Store(dst)
}

func (m *Map) Load() map[string]*Value {
	return m.values.Load().(map[string]*Value)
}

func (m *Map) Exist(key string) bool {
	_, ok := m.Load()[keyNamed(key)]
	return ok
}

func (m *Map) Get(key string) *Value {
	v, ok := m.Load()[keyNamed(key)]
	if ok {
		return v
	}

	return &Value{}
}

func (m *Map) Keys() []string {
	values := m.Load()
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}

	return keys
}

func keyNamed(key string) string {
	return strings.ToLower(key)
}
