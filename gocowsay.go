package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// main entry point
func main()  {
  var output []rune


  info, _ := os.Stdin.Stat()

  if info.Mode()&os.ModeCharDevice != 0 {
    fmt.Println("The command is intended to work with pipes.")
    fmt.Println("Usage Ex: fortune | gocowsay")
    return
  }

  reader := bufio.NewReader(os.Stdin)

  for {
    input, _, err := reader.ReadRune()
    if err != nil && err == io.EOF {
      break
    }

    output = append(output, input)
  }

  for j:=0; j<len(output); j++{
    fmt.Printf("%c", output[j])
    
  }
    
}
