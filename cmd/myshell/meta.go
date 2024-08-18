package main

import (
	"fmt"
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
			//fmt.Printf("err: %+v", err)
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			command[entry.Name()] = fmt.Sprintf("%s/%s", path, entry.Name())
		}
	}
	_meta.command = command
}
