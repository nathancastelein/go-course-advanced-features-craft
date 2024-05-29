package generics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapKeys(t *testing.T) {
	// Arrange
	m := NewMap[string, int]()
	m.Store("hello", 1)
	m.Store("world", 2)

	// Act
	keys := m.Keys()

	// Assert
	require.Len(t, keys, 2)
	require.Equal(t, []string{"hello", "world"}, keys)
}

func TestMapValues(t *testing.T) {
	// Arrange
	m := NewMap[string, int]()
	m.Store("hello", 1)
	m.Store("world", 2)

	// Act
	keys := m.Values()

	// Assert
	require.Len(t, keys, 2)
	require.Equal(t, []int{1, 2}, keys)
}
