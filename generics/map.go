package generics

type Map[Key comparable, Value any] struct {
	store map[Key]Value
}

func NewMap[Key comparable, Value any]() *Map[Key, Value] {
	return &Map[Key, Value]{
		store: map[Key]Value{},
	}
}

func (m *Map[Key, Value]) Store(key Key, value Value) {
	m.store[key] = value
}

func (m Map[Key, Value]) Size() int {
	return len(m.store)
}

func (m Map[Key, Value]) Keys() []Key {
	keys := make([]Key, len(m.store))
	var i int
	for key := range m.store {
		keys[i] = key
		i++
	}

	return keys
}

func (m Map[Key, Value]) Values() []Value {
	values := make([]Value, len(m.store))
	var i int
	for _, value := range m.store {
		values[i] = value
		i++
	}

	return values
}
