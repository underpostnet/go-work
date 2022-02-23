

package main


import (
  "fmt"
  "encoding/json"
)

/*

bool, string,
int,  int8,  int16,  int32,  int64,
uint, uint8, uint16, uint32, uint64, uintptr,
byte, // alias for uint8
rune, // alias for int32
float32, float64,
complex64, complex128

*/

type ObjTest struct {
  a string
  b int
  c bool
  d int
  e string
}


func main(){


  fmt.Println(" test")

  stringJSON := `[
    {
      "a": "A",
      "b": 1,
      "c": true,
      "d": null,
      "e": null
    },
    {
      "a": "B",
      "b": 2,
      "c": false,
      "d": null,
      "e": null
    }
  ]
`;

  var objs []ObjTest
  json.Unmarshal([]byte(stringJSON), &objs)
  fmt.Printf("ObjTest : %+v", objs)



}
