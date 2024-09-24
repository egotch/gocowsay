package main

import (
	"flag"

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

  figPtr := flag.String("f", "rnd", "select which animal figure to have talk back cow, stego, or rnd!")

  flag.Parse()

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
  animal := ascii.GetAnimal(*figPtr)

  fmt.Println(balloon)
  fmt.Println(animal)
  fmt.Println()

}
