// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func main() {
  var p *int32 = new(int32)
  var i int32
  
  var p2 *int32 = &i
  
  fmt.Println("p is", p)
  fmt.Println("i is", i)
  fmt.Println("p2 is", p2)
}
