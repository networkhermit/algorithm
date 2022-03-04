package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("When’s Weekend?")

    today := time.Now().Weekday()

    switch today {
    case time.Sunday:
        fmt.Println("Today.")
    case time.Saturday:
        fmt.Println("Today.")
    case time.Friday:
        fmt.Println("Tomorrow.")
    case time.Thursday:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
}
