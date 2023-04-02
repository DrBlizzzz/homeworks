package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDir(t *testing.T) {
	t.Run("correct_environment", func(t *testing.T) {
		results := map[string]EnvValue{
			"FOO":   {"   foo\nwith new line", false},
			"HELLO": {"\"hello\"", false},
			"UNSET": {"", true},
			"BAR":   {"bar", false},
			"EMPTY": {"", false},
		}
		env, _ := ReadDir("testdata/env")
		for k, v := range env {
			assert.Equal(t, results[k].Value, v.Value)
		}
	})
}
