package main

import "fmt"

func main()  {
  fmt.Printf("%v\n%v\n%v\n", func () (arr [3]int, slice []int, m map[int]string)  {
    //noraml arrays (static size)
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3

    //slices
    //not great growth.... it creates a new slice to copy everythign into,
    //slicing into smaller pieces is extremly efficent and wastes zero space
    //
    //important to note that when you slice a slice it still points to the
    //original place in memory.
    slice[00] = 01
    slice[01] = 10
    slice[10] = 11
    slice[11] = 00

    //maps are like python (hashmaps)
    m[0] = "Winter"
    m[1] = "Spring"
    m[2] = "Summer"
    m[3] = "Fall"
    return
  })
}
