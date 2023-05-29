package littlelib

import (
  "fmt"
  "example.com/skeleton/tinylib/sublib"
)

func Hi() {
  fmt.Println("Hi from tinylib!")
  fmt.Println("  Notice the file name doesn't matter.")
  fmt.Println("  The package name 'littlelib' does matter, it's the default")
  fmt.Println("  name in the importing file, if not overriden during import")
  fmt.Println("")

  tinysub.SayHi()
}

