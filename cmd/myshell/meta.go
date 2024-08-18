package main

import (
	"flag"
	"fmt"
	"strings"
)

var _meta = meta{}

type meta struct {
	paths []string
}

func initMeta() {
	var path string
	flag.StringVar(&path, "path", "", "path")
	flag.Parse()

	paths := strings.Split(path, ":")
	_meta.paths = paths
	fmt.Printf("%q", paths)
}
