package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	initMeta()

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
