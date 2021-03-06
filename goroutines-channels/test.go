
/*

new(int)        -->  NEW(*int)
new(Point)      -->  NEW(*Point)
new(chan int)   -->  NEW(*chan int)
make([]int, 10) -->  NEW([]int, 10)

make(Point)  // Illegal
make(int)    // Illegal


Things you can do with make that you can't do any other way:

  . Create a channel
  . Create a map with space preallocated
  . Create a slice with space preallocated or with len != cap


"make" function allocates and initializes an object of type slice, map, or chan only.
Like new, the first argument is a type. But, it can also take a second argument,
the size. Unlike new, make’s return type is the same as the type of its argument,
not a pointer to it. And the allocated value is initialized (not set to zero
value like in new). The reason is that:

 slice, map and chan are data structures.


They need to be initialized, otherwise they won't be usable. This is the reason
new() and make() need to be different.

*/

package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)


func RandomInt(min int, max int) int {
    return rand.Intn((max + 1) - min) + min
}

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {

				// time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
				
				randomWait := RandomInt(1, 5)
					fmt.Println(" pre random wait -> ")
					fmt.Println(randomWait)
				time.Sleep(time.Duration(randomWait) * time.Second)
				ch <- "Goroutine : " + strconv.Itoa(i) + " randomWait: " + strconv.Itoa(randomWait)

			}
		}(i)
	}

	fmt.Println(" async ")

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-ch)
	}

	fmt.Println(" end ")
}





























/**/
