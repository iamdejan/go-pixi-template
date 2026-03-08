package message

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	names := []string{
		"alpha",
		"beta",
		"gamma",
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			expected := fmt.Sprintf("Hello, %s", name)
			assert.Equal(t, expected, Hello(name))
		})
	}
}
