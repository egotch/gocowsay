package helpers_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/egotch/gocowsay/helpers"
)

var testLines = []string{
  "this\t is a test",
  "if\thandled correctly, should not have any \ttabs left",
}


func Test_tabsToSpaces(t *testing.T) {


  // positive test
  result := helpers.TabsToSpaces(testLines)
  for _, l := range result {
    if strings.Contains(l, "\t") {
      t.Error("incorrect result: line contains tabs. line ->", l)
    }
  }
  // negative test (control)
  result = testLines
  for _, l := range result {
    if ! strings.Contains(l, "\t"){
      t.Error("Control test fail: original line should contain tabs. line ->", l)
    }
  }


}

func Test_calcMaxWidth(t *testing.T) {
  
  scrubbedText := helpers.TabsToSpaces(testLines)

  // positive test
  result := helpers.CalcMaxWidth(scrubbedText) 
  if result != 58 {
    t.Error("incorrect result: line length = ", result)
  }

  // negative test
  if result == 0 {
    t.Error("fail, result cannot be zero")
  }
}

func Test_normalizeLineLength(t *testing.T)  {

  scrubbedText := helpers.TabsToSpaces(testLines)
  maxLen := helpers.CalcMaxWidth(scrubbedText)
  result := helpers.NormalizeLineLen(scrubbedText, maxLen)


  // postive test - test that all lines are the same length
  currLen := utf8.RuneCountInString(result[0])
  for _, l := range result {
    if utf8.RuneCountInString(l) != currLen{
      t.Error("incorrect result: line length doesnt match max line length", l)
    }
  }
  
}
