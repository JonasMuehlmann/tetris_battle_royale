package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, 1, 2, "Oh no!")
}
