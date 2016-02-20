package main

import (
  "os"
  "os/exec"
)

func RunCommand(command string, arguments ...string) {
  cmd := exec.Command(command, arguments...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}
