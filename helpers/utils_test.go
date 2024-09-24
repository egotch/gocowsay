package helpers_test

import (
	"strings"
	"testing"

	"github.com/egotch/gocowsay/helpers"
)

func Test_tabsToSpaces(t *testing.T) {

  testLines := []string{
  "this\t is a test",
  "if\thandled correctly, should not have any \ttabs left",
  }

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

//TODO: add unit test for CalcMaxWidth

//TODO: add unit test for NormalizeLineLengh
