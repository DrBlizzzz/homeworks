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

	// добавил тест на проверку наличия = в имени
	t.Run("check files with = in name", func(t *testing.T) {
		results := map[string]EnvValue{
			"TEST2": {"TEST2", false},
		}
		env, _ := ReadDir("testdata/myenv")
		for k, v := range env {
			assert.Equal(t, results[k].Value, v.Value)
		}
	})
}
