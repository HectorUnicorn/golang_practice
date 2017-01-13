package main

import (
        "fmt"
        "math"
        "time"
)

func main() {
        fmt.Println(Sqrt(2))
        fmt.Println(Sqrt(-2))
}

type MyError struct {
        When time.Time
        What string
}

func (e *MyError) Error() string {
        return fmt.Sprintf("at %v, %s",
                e.When, e.What)
}

func run() error {
        return &MyError{
                time.Now(),
                "it didn't work",
        }
}

func Sqrt(x float64) (float64, error) {
        if x >= 0 {
                return math.Sqrt(x), nil
        } else {
                return 0, &MyError{time.Now(), "x is < 0"}
        }
        //return 0, nil
}
