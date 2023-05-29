package main

import (
  "fmt"
  "example.com/skeleton/tinylib"
)

func main() {
  fmt.Println("Hi from main!");
  fmt.Println("  This organization demonstrates how a separate library module")
  fmt.Println("  and binary module can be developed in the same code repo.")
  fmt.Println("")

  littlelib.Hi()
}

