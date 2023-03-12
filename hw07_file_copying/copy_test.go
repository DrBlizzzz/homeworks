package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("offset0_limit0", func(t *testing.T) {
		Copy("testdata/input.txt", "out.txt", 0, 0)
		result, _ := os.ReadFile("out.txt")
		target, _ := os.ReadFile("testdata/out_offset0_limit0.txt")
		require.Equal(t, target, result)
	})

	t.Run("offset0_limit10", func(t *testing.T) {
		Copy("testdata/input.txt", "out.txt", 0, 10)
		result, _ := os.ReadFile("out.txt")
		target, _ := os.ReadFile("testdata/out_offset0_limit10.txt")
		require.Equal(t, target, result)
	})

	t.Run("offset0_limit1000", func(t *testing.T) {
		Copy("testdata/input.txt", "out.txt", 0, 1000)
		result, _ := os.ReadFile("out.txt")
		target, _ := os.ReadFile("testdata/out_offset0_limit1000.txt")
		require.Equal(t, target, result)
	})

	t.Run("offset0_limit10000", func(t *testing.T) {
		Copy("testdata/input.txt", "out.txt", 0, 10000)
		result, _ := os.ReadFile("out.txt")
		target, _ := os.ReadFile("testdata/out_offset0_limit0.txt")
		require.Equal(t, target, result)
	})

	t.Run("offset100_limit1000", func(t *testing.T) {
		Copy("testdata/input.txt", "out.txt", 100, 1000)
		result, _ := os.ReadFile("out.txt")
		target, _ := os.ReadFile("testdata/out_offset100_limit1000.txt")
		require.Equal(t, target, result)
	})

	t.Run("offset6000_limit1000", func(t *testing.T) {
		Copy("testdata/input.txt", "out.txt", 6000, 1000)
		result, _ := os.ReadFile("out.txt")
		target, _ := os.ReadFile("testdata/out_offset6000_limit1000.txt")
		require.Equal(t, target, result)
	})
}
