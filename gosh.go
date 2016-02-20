package main

import (
	"fmt"
)

func main() {
	config := LoadConfig()

	fmt.Printf(LocateCommand("ls", config.Path))
}
