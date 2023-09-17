// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func main() {
  var myArr [3]int32 = [3]int32{1, 2, 3}
  var myArr1 [3]int32 = [...]int32{1, 2, 3} // three dots operator, automatically detect number
  
//   myArr[0] = 0
//   myArr[1] = 1
//   myArr[2] = 2
  
  fmt.Println(myArr)
  fmt.Println(myArr[0])
  fmt.Println(myArr[1])
  fmt.Println(myArr[2])
  fmt.Println(myArr[1:2])
  fmt.Println(myArr[1:3])
    fmt.Println(&myArr)
    fmt.Println(&myArr[0])
  
  fmt.Println(myArr1)
  
    var mySlice []int32 = []int32{4, 5, 6}
    fmt.Println(mySlice)
    mySlice = append(mySlice, 7)
    fmt.Println(mySlice)
    
    var mySlice2 = []int32{8, 9, 10}
    mySlice = append(mySlice, mySlice2...) // spread operator
    
    fmt.Println(mySlice)
    
    var mySlice3 = make([]int32, 3, 8) // type, size and capacity
    fmt.Println(mySlice3)
    
    var myMap map[string]uint32 = make(map[string]uint32)
    var myMap2 = map[string]uint32{"Pushpa": 24, "Pushpa2": 23}
    fmt.Println(myMap)
    fmt.Println(myMap2["Pushpa"])
    
    if val, ok := myMap2["NA_KEY"]; ok {
        fmt.Println("val is " + string(val))
    } else {
        fmt.Println("key does not exist")
    }
    
    delete(myMap2, "Pushpa")
    fmt.Println(myMap2)
    
    for k, v := range myMap2 {
        fmt.Printf("key: %v, val: %v\n", k, v)
    }
    
    for i, v := range mySlice2 {
        fmt.Printf("index: %v, val: %v\n", i, v)
    }
    
    // strings are immutable in go, use the built-in strings package to make a string builder
}
