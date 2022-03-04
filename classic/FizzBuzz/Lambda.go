package main

import (
    "fmt"
    "strconv"
)

func fizzBuzz(n int) string {
    var word string

    if n % 3 == 0 {
        word += "Fizz"
    }
    if n % 5 == 0 {
        word += "Buzz"
    }

    if len(word) == 0 {
        word = strconv.Itoa(n)
    }

    return word
}

func main() {
    for i := 1; i <= 100; i++ {
        fmt.Println(fizzBuzz(i))
    }
}
