package main

import (
	"bufio"
	"fmt"
	"os"
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
	}
}

func handleCommand(command string) {
	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
}
