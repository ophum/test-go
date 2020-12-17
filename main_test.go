package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	msg := HelloWorld()
	assert.Equal(t, "Hello world", msg)
}
