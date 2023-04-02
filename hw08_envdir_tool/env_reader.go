package main

import (
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	varFiles, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	env := make(Environment)
	for _, entry := range varFiles {
		entryName := entry.Name()
		buffer, err := os.ReadFile(path.Join(dir, entryName))
		if err != nil {
			return nil, err
		}
		if len(buffer) == 0 {
			env[entryName] = EnvValue{"", true}
			break
		}
		for index, splitByte := range buffer {
			// первый раз встретили \n
			if splitByte == byte(10) {
				buffer = buffer[:index]
				break
			}
			// замена нуль терминала на \n
			if splitByte == byte(0) {
				buffer[index] = byte(10)
			}
		}
		env[entryName] = EnvValue{strings.TrimRight(string(buffer), " "), false}
	}
	return env, nil
}
