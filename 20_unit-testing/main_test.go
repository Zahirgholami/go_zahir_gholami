package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Greeting() string {
	return "Hello, World!"
}

func TestGreeting(t *testing.T) {
	expected := "Hello, World!"
	actual := Greeting()
	assert.Equal(t, expected, actual, "Greeting message should be 'Hello, World!'")
}
