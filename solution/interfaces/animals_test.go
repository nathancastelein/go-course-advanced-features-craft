package interfaces

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeSomeNoise(t *testing.T) {
	// Arrange
	orig := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	defer func() {
		os.Stdout = orig
	}()

	// Act
	MakeSomeNoise([]Animal{
		Cat{},
		Dog{},
		Cat{},
	})

	// Assert
	expectedOutput := strings.Join([]string{(Cat{}).Speak(), (Dog{}).Speak(), (Cat{}).Speak()}, "\n") + "\n"
	w.Close()
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	require.Equal(t, expectedOutput, string(out))
}
