package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	config := LoadConfig()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("exit\n")
			os.Exit(0)
		}
		tokens := strings.Split(strings.TrimSpace(text), " ")
		command := tokens[0]
		arguments := tokens[1:]
		if len(command) > 0 {
			command = LocateCommand(command, config.Path)
			RunCommand(command, arguments...)
		}
	}
}
