package main

import (
	"fmt"
	"os"
	"os/exec"
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
		reloadMeta()
		handleType(strs[1])
	default:
		reloadMeta()
		path, isPath := _meta.command[strs[0]]
		if !isPath {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
			return
		}
		out, err := exec.Command(path, strs[1:]...).Output()
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stdout, "%s", string(out))
	}

}

func handleExit(code int) {
	os.Exit(code)
}

func handleEcho(str string) {
	fmt.Fprintf(os.Stdout, "%s\n", str)
}

func handleType(_type string) {
	switch _type {
	case "exit", "echo", "type":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", _type)
	default:
		path, isPath := _meta.command[_type]
		if !isPath {
			fmt.Fprintf(os.Stdout, "%s: not found\n", _type)
			return
		}
		fmt.Fprintf(os.Stdout, "%s is %s\n", _type, path)
	}
}
