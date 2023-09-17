// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

// small letter starting means that it is private to the package, big letter starting means it is public with the package

type owner struct {
    name string
}

// type gasEngine struct {
//     mpg uint8
//     gallons uint8
//     ownerInfo owner
// }

type GasEngine struct {
    mpg uint8
    gallons uint8
    owner
    int
}

func (e GasEngine) milesLeft() uint8 {
    return e.gallons * e.mpg
}

type ElectricEngine struct {
    mpkwh uint8
    kwh uint8
}

func (e ElectricEngine) milesLeft() uint8 {
    return e.mpkwh * e.kwh
}

type Engine interface {
    milesLeft() uint8
}

// func canMakeIt(e GasEngine, miles uint8) {
//     if miles <= e.milesLeft(){
//         fmt.Println("You can make it.")
//     } else {
//         fmt.Println("No point trying, just give up.")
//     }
// }

// making generic for both the engines using interfaces
func canMakeIt(e Engine, miles uint8) {
    if miles <= e.milesLeft(){
        fmt.Println("You can make it.")
    } else {
        fmt.Println("No point trying, just give up.")
    }
}

func main() {
  var myGasEngine GasEngine = GasEngine{25, 32, owner{"Pushpa"}, 10}
//   fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
  
  fmt.Println(myGasEngine.mpg, myGasEngine.gallons, myGasEngine.name, myGasEngine.int)
  fmt.Println(myGasEngine.milesLeft())
  canMakeIt(myGasEngine, 200)
  
  var myElectricEngine ElectricEngine = ElectricEngine{25, 32}
  
  fmt.Println(myElectricEngine.mpkwh, myElectricEngine.kwh)
  fmt.Println(myElectricEngine.milesLeft())
  canMakeIt(myElectricEngine, 20)
  
  var genericEngine Engine = myGasEngine
  canMakeIt(genericEngine, 200)
}
