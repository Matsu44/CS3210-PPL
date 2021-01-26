// CS 3210 - Principles of Programming Languages - Spring 2020
// Homework 11 - Complete the TO-DO's embedded in the code
// Author: Xiaosong Wang

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
    "sort"
)

var wg sync.WaitGroup

// global constants definitions
const N       = 100000
const MAX_INT = 10000

// this function sorts a slice using the awesome insertion sort algorithm!
func sorts(slice []int) {
  // comment the line below when timing without using goroutines
  defer wg.Done()
  for i, _ := range slice {
    j := i
    k := i - 1
    for k >= 0 {
      if slice[j] < slice[k] {
        temp := slice[j]
        slice[j] = slice[k]
        slice[k] = temp
        j = k
        k--
      } else {
        break
      }
    }
  }
}

// this function merges two slices together
func merge(left []int, right []int) {
  var data = make([]int, len(left) + len(right))
  i := 0
  j := 0
  k := 0
  for ; i < len(left) && j < len(right); k++ {
    if left[i] < right[j] {
      data[k] = left[i]
      i++
    } else {
      data[k] = right[j]
      j++
    }
  }
  for ; i < len(left); i++ {
    data[k] = left[i]
    k++
  }
  for ; j < len(right); j++ {
    data[k] = right[j]
    k++
  }
  k = 0
  for i := 0; i < len(left); i++ {
    left[i] = data[k]
    k++
  }
  for j := 0; j < len(right); j++ {
    right[j] = data[k]
    k++
  }
}

func main() {

  // sets the seed of the random generator
  rand.Seed(0)

  // TODO: generate an array of N = 100K random integers in [0,MAX_INT)
  var slice = make([]int, N)
  i := 0
  for i < N {
    slice[i] = rand.Intn(MAX_INT)
    i++
  }

  // begins timing the procedure
  start := time.Now()

  // TODO: sort the array without using goroutines
  sort.Ints(slice)

  // TODO: sort the array using 2 goroutines
  // begin by creating 2 equally sized slices
  // then pass each slice to a goroutine based on the "sort" function
  // wait for the goroutines to end before merging the slices
  left := slice[:MAX_INT/2]
  right := slice[:MAX_INT/2]
  wg.Add(2)
  go sorts(left)
  go sorts(right)
  wg.Wait()
  merge(left, right)

  // ends timining and shows how long it took
  fmt.Println(time.Since(start))
}
