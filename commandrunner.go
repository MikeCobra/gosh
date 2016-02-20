package main

import (
  "os"
  "os/exec"
)

func RunCommand(command string) {
  cmd := exec.Command(command)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}
