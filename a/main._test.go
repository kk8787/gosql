package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJson(t *testing.T) {
	// json := `{"name": "John", "age": 30, "city": "New York"}`
	// parsed := parseJson(json)

	// // asert
	// assert.Equal(t, parsed, map[string]any{"name": "John", "age": 30, "city": "New York"})

	json := `{"name": "John"}`
	expected := []Token{
		Token("{"),
		Token("name"),
		Token(":"),
		Token("John"),
		Token("}"),
	}
	parsed := lex(json)

	// check length
	assert.Equal(t, len(expected), len(parsed))

	// check tokens
	assert.Equal(t, expected, parsed)
}
