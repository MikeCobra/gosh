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
		command := strings.Split(strings.TrimSpace(text), " ")[0]
		if len(command) > 0 {
			fmt.Printf("%s\n", LocateCommand(command, config.Path))
		}
	}
}
