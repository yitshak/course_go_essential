package main

import (
  "fmt"
  "strings"
)

func main() {
  test := "cehck check this out what out this is foo bar bar foo for chum"

  words := strings.Fields(test)

  word_map := map[string]int{}

  for _, word := range words {
    word_map[word] = word_map[word]+1
  }

  for word, count := range word_map{
    fmt.Printf("%v :%v\n", word, count)  
  }
}