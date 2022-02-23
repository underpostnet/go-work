

package main

import (
  "fmt"
  "test/dir1"
  "test/dir2"
  "test/dir3"
)


func main() {

    fmt.Println("test")
    fmt.Println(dir1.SayHello())
    fmt.Println(dir2.SayHello())
    fmt.Println(dir3.SayHello())

}
