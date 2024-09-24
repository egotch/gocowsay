package helpers

import (
	"fmt"
	"strings"
)

// BuildBalloon is a public function
// taks in a slice of strings of max width maxwidth
// prepends/appends margins on the first and last line, and at the start/end of each line
// and returns a string with the the contents of the balloon
func BuildBalloon(lines []string, maxwidth int) string  {

  // var declarations
  var boarders []string
  count := len(lines)
  var ret []string
  var innards []string

  // slice of boarders
  boarders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

  // build top of bubble
  topBtm := " " + strings.Repeat("-", maxwidth+2)

  // build innards of bubble
  // single line
  if count == 1 {
    s := fmt.Sprintf("%s %s %s", boarders[5], lines[0], boarders[6])
    innards = append(innards, s)
    
  } else {
    // multiple lines

    // first line
    s := fmt.Sprintf("%s %s %s", boarders[0], lines[0], boarders[1])
    innards = append(innards, s)

    // middle lines
    for i := 1; i < count-1; i++ {
      s = fmt.Sprintf("%s %s %s", boarders[4], lines[i], boarders[4])
      innards = append(innards, s)
    }

    // last line
    s = fmt.Sprintf("%s %s %s", boarders[2], lines[len(lines)-1], boarders[3])
    innards = append(innards, s)
    
  }

  // build the return
  ret = append(ret, topBtm)
  ret = append(ret, innards...)
  ret = append(ret, topBtm)
  return strings.Join(ret, "\n")
}

