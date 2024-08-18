package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleCommand(command string) {
	strs := strings.Split(command, " ")

	switch strs[0] {
	case "exit":
		code, _ := strconv.Atoi(strs[1])
		handleExit(code)
	case "echo":
		handleEcho(strings.Join(strs[1:], " "))
	case "type":
		handleType(strs[1])
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}

}

func handleExit(code int) {
	os.Exit(code)
}

func handleEcho(str string) {
	fmt.Fprintf(os.Stdout, "%s\n", str)
}

func handleType(_type string) {
	path, isPath := _meta.command[_type]
	if isPath {
		fmt.Fprintf(os.Stdout, "%s is %s\n", _type, path)
		return
	}

	switch _type {
	case "exit", "echo", "type":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", _type)
	default:
		fmt.Fprintf(os.Stdout, "%s: not found\n", _type)
	}
}
