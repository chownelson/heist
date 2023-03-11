package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)
//first
  if eludedGuards >= 50{
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else{
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
//second
  openedVault := rand.Intn(100)
  if openedVault >= 70 && isHeistOn{
    fmt.Println("Grab and GO!")
  } else if isHeistOn{
    isHeistOn = false
    fmt.Println("The vault can't be opened.")
  }
//third
  if isHeistOn{
    leftSafely := rand.Intn(5)
    switch leftSafely{
      case 0:
        isHeistOn = false
        fmt.Println("Heist has failed. (LOSER!!!!!)")
      case 1:
        isHeistOn = false
        fmt.Println("Turns out vault doors don't open from the inside...")
      case 2:
        isHeistOn = false
        fmt.Println("Your partner betrayed you")
      case 3:
        isHeistOn = false
        fmt.Println("Someone triggered the silent alarm, should've seen that coming...")
      default:
        fmt.Println("Start the getaway car!")
        if isHeistOn{
          amtStolen := 10000 + rand.Intn(1000000)
          fmt.Printf("You stole $%v !!!", amtStolen)
        }
    }

}
}
