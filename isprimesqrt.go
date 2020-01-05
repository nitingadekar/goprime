package main

import (
	"fmt"
    "math"
    "strconv"
)


func IsPrimeSqrt(value int) bool {
    for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
        if value%i == 0 {
		//fmt.Println("Number is not prime",value )
            return false
        }
    }
    //fmt.Println("Number is prime", value)
    return value > 1
}

func main(){
	to :=  379
	t := strconv.Itoa(to)
        fmt.Println(t)
	fmt.Println(len(t))
	if IsPrimeSqrt( to ) {
		fmt.Printf("true")
	} else {
		fmt.Printf("flase")
	}
}
