package main

import (
  "github.com/egotch/gocowsay/ascii"
  "github.com/egotch/gocowsay/helpers"

	"bufio"
	"fmt"
	"io"
	"os"
)

// main entry point
func main()  {

  var lines []string

  info, _ := os.Stdin.Stat()

  if info.Mode()&os.ModeCharDevice != 0 {
    fmt.Println("The command is intended to work with pipes.")
    fmt.Println("Usage Ex: fortune | gocowsay")
    return
  }

  // fetch and iterate over the input
  reader := bufio.NewReader(os.Stdin)

  for {
    line, _, err := reader.ReadLine()
    if err != nil && err == io.EOF {
      break
    }

    lines = append(lines, string(line))
  }

    
  // build the bubble and cow
  lines = helpers.TabsToSpaces(lines)
  maxWidth := helpers.CalcMaxWidth(lines)
  normLines := helpers.NormalizeLineLen(lines, maxWidth)
  balloon := helpers.BuildBalloon(normLines, maxWidth)
  cow := ascii.Cow

  fmt.Println(balloon)
  fmt.Println(cow)
  fmt.Println()

}
