package main

import (
    "fmt"
    "math"
)

func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
	    fmt.Println("value is %v", i)
        if value%i == 0 {
	    fmt.Println("This is false")
            return false

        }
    }
    fmt.Println("This is true")
    return value > 1

}




func main() {
	IsPrime(33)
}
