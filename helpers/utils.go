package helpers

import (
	"strings"
	"unicode/utf8"
)

// tabsToSpaces converts all tabs found in the strings
// converts a tab to 4 spaces to prevent misalignment in the "bubble"
func TabsToSpaces(lines []string) []string  {
  var ret []string 
  for _, v := range lines {
    v = strings.Replace(v, "\t", "    ", -1)
    ret = append(ret, v)
  }

  return ret
}


// calcMaxWidth determines the longest line
// returns int -> length of longest line
func CalcMaxWidth(lines []string) int  {
  var maxLen int 

  maxLen = 1
  for _, v := range lines{
    if utf8.RuneCountInString(v) > maxLen {
      maxLen = utf8.RuneCountInString(v)
    } 
  }

  return maxLen
}


// normalizeLineLen: takes a slice of strings
// appends whitespace to the end of each line to make them all the same length
func NormalizeLineLen(lines []string, maxLen int) []string  {
 
  var ret []string

  // iterate over lines
  for _, v := range lines {
    normLen := maxLen - utf8.RuneCountInString(v)
    s := v

    if normLen > 0 {
      s = v + strings.Repeat(" ", normLen)
    } 

    ret = append(ret, s)
  }

  return ret
}
