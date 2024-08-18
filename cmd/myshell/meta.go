package main

import (
	"os"
	"strings"
)

var _meta = meta{}

type meta struct {
	paths   []string
	command map[string]string
}

func initMeta() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	_meta.paths = paths
	command := make(map[string]string)

	for _, path := range paths {
		entries, err := os.ReadDir(path)
		if err != nil {
			os.Exit(-1)
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			command[entry.Name()] = path
		}
	}
	_meta.command = command
}
