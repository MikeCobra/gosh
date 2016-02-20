package main

import (
  "io/ioutil"
)


func LocateCommand(command string, path []string) (commandFullPath string) {
  for _, directory := range path {
    files, err := ioutil.ReadDir(directory)
    if err != nil {
      panic(err)
    }
    for _, file := range files {
      if file.Name() == command{
        return directory + "/" + command
      }
    }
  }
  return "Gone bad!"
}
