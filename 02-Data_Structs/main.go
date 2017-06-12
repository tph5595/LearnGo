package main

import "fmt"

func main()  {
  fmt.Printf("%v", func () (arr []int)  {
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    return
  })
}
