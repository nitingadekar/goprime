package main

import (
	"fmt"
	"strconv"
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




	//fmt.Println(string(t[:len(t)-1]))              // ASCII only


func main() {
	prime := 6789
	fmt.Printf("%T\n", prime)
	t := strconv.Itoa(prime)
	fmt.Printf("%T\n", t)
	fmt.Println(string(t[1:]))              // ASCII only
	fmt.Println(int(t))


    //fmt.Println(string([]rune("Hello, pa")[1])) // UTF-8
    //fmt.Println(string([]rune("Hello, ma")[8])) // UTF-8

  //  return IsPrime(6789)

}
