// CS 3210 - Principles of Programming Languages - Spring 2020
// Homework 10 - Complete the TO-DO's embedded in the code
// Author: Xiaosong Wang

package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "io"
)

const FILE_NAME = "common.txt"
const NUMBER_WORDS = 1000

// TODO: finish the function implementation
func loadWords(fileName string) []string {

  // TODO: open the file using fileName parameter
  txtFile, _ := os.Open(FILE_NAME)

  // TODO: use a bufio.NewReader
  reader := bufio.NewReader(txtFile)

  // words is a slice with zero length and NUMBER_WORDS capacity
  words := make([]string, 0, NUMBER_WORDS)

  // TODO: use a for loop to read all words from the file into the words slice (hint: use the append function)
 for{
   line, error := reader.ReadString('\n')

   if error == io.EOF {
     break
   }

   line = strings.TrimSuffix(line, "\n")
   words = append(words, line)
 }

  // words is returned
  return words
}

// TODO: finish the function implementation
func hasWord(words []string, search string) bool {

    // TODO: use a for loop to iterate over the words slice
    // if the search word is found in the slice, return true
    for _, str := range words{
      if search == str {
        return true
      }
    }
    // return false now because the search word wasn't found
    return false
}

func main() {
  words := loadWords(FILE_NAME)

  if hasWord(words, "fox") {
    fmt.Println("fox was found!")
  } else {
    fmt.Println("fox was not found!")
  }

  if hasWord(words, "grass") {
    fmt.Println("grass was found!")
  } else {
    fmt.Println("grass was not found!")
  }
}
