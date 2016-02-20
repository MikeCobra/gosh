package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	config := LoadConfig()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		command := strings.Split(strings.TrimSpace(text), " ")[0]
		fmt.Printf("%s\n", LocateCommand(command, config.Path))
	}
}
