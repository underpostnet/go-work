

package main


import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "log" // log.Fatal(err)
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

func RecurseJsonMap(dat map[string]interface{}) {
    fmt.Println("object {")
    for key, value := range dat {
        fmt.Print("key: " + key + " ")
    // print array properties
    arr, ok := value.([]interface{})
    if ok {
        fmt.Println("value: array [")
        for _, arrVal := range arr {
            // recurse subobjects in the array
            subobj, ok := arrVal.(map[string]interface{})
            if ok {
                RecurseJsonMap(subobj)
            } else {
                // print other values
                fmt.Printf("value: %+v\n", arrVal)
            }
        }
        fmt.Println("]")
    }

    // recurse subobjects
    subobj, ok := value.(map[string]interface{})
    if ok {
        RecurseJsonMap(subobj)
    } else {
            // print other values
            fmt.Printf("value: %+v\n" ,value)
        }
    }
    fmt.Println("}")
}



type ObjTest struct {
  a string `json:"a"`
  b int
  c bool
  d int
  e string
}


func main(){


  fmt.Println(" test")

  // Read entire file content, giving us little control but
  // making it very simple. No need to close the file.
  content, err := ioutil.ReadFile("./data/test.json")
  if err != nil {
      log.Fatal(err)
  }

  // Convert []byte to string and print to screen
  stringJSON := string(content)
  fmt.Printf(stringJSON)

  var objs []ObjTest
  json.Unmarshal([]byte(stringJSON), &objs)

  /*

  https://pkg.go.dev/fmt
  %v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
  %#v	a Go-syntax representation of the value
  %T	a Go-syntax representation of the type of the value
  %%	a literal percent sign; consumes no value

  */


  // we are parsing it in a generic fashion
  byt := []byte((`{ "data":`+stringJSON+`}`))
  var dat map[string]interface{}
  if err := json.Unmarshal(byt, &dat); err != nil {
      panic(err)
  }
  RecurseJsonMap(dat)



}
