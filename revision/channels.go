// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func main() {
    /*
        channels -> 1. hold data, 2. thread safe, 3. listen for data
        * when you write to an unbuffered channel, it is blocked until it is received, if none receive it, deadlock happens
        * By default channels are unbuffered, which states that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) which are ready to receive the sent value.
        * Buffered channel are blocked only when the buffer is full. Similarly receiving from a buffered channel are blocked only when the buffer will be empty.
        * channles + goroutines = ♥
        * make sure to close the channel, or face the deadlock
    */
    
    var ch = make(chan int) // unbuffered channel
    ch <- 1
    
    /* ------------- Deadlock ------------- */
    var  i = <- ch
    fmt.Println(i)
    /* ------------- Deadlock ------------- */
}
