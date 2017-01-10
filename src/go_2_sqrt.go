// the go customized sqrt function
package main

import (
        "fmt"
        "math"
        "flag"
)

var number float64

func init() {
        flag.Float64Var(&number, "n", 2.0, "input a number to sqrt")
        flag.Parse()
}

func Sqrt(x float64) float64 {
        z := float64(x)
        s := float64(0)
        for {
                z = z - (z*z - x)/(2*z)
                if math.Abs(s-z) < 1e-10 {
                        break;
                }
                s = z
        }
        return s
}

func main() {
        fmt.Println(Sqrt(number))
        //fmt.Println(Sqrt(2))
}