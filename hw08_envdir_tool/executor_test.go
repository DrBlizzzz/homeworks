package main

import (
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("check_executor", func(t *testing.T) {
		RunCmd([]string{"/bin/bash", "-c", "pwd"}, nil)
	})
}
