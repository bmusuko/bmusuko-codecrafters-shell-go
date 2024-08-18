package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	for {
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(-1)
		}
		command = strings.TrimSuffix(command, "\n")
		handleCommand(command)

		fmt.Fprint(os.Stdout, "$ ")
	}
}

func handleCommand(command string) {
	strs := strings.Split(command, " ")

	switch strs[0] {
	case "exit":
		code, _ := strconv.Atoi(strs[1])
		handleExit(code)
	case "echo":
		handleEcho(strings.Join(strs[1:], " "))
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
