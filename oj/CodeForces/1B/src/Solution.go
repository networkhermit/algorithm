package main

import (
    "fmt"
    "io"
    "unicode"
)

func process(coordinate string) {
    column := 0
    row := 0

    length := len(coordinate)

    k := 1
    if rune(coordinate[0]) == 'R' && !unicode.IsLetter(rune(coordinate[1])) {
        for k < length && unicode.IsNumber(rune(coordinate[k])) {
            k++
        }
    }

    if k == 1 || k == length {
        k = 0
        for unicode.IsLetter(rune(coordinate[k])) {
            column = column * 26 + int(coordinate[k] - 64)
            k++
        }
        for k < length {
            row = row * 10 + int(coordinate[k] - 48)
            k++
        }

        fmt.Printf("R%dC%d\n", row, column)
    } else {
        k = 1
        for unicode.IsNumber(rune(coordinate[k])) {
            row = row * 10 + int(coordinate[k] - 48)
            k++
        }
        k++
        for k < length {
            column = column * 10 + int(coordinate[k] - 48)
            k++
        }

        s := make([]string, 0)

        var r int
        for column != 0 {
            r = column % 26
            column /= 26
            if r == 0 {
                s = append(s, "Z")
                column--
            } else {
                s = append(s, string(r + 64))
            }
        }

        for i := len(s) - 1; i >= 0; i-- {
            fmt.Print(s[i])
        }
        fmt.Println(row)
    }
}

func main() {
    var testCases int

    var coordinate string

    for {
        _, err := fmt.Scan(&testCases)
        if err == io.EOF {
            break
        }
        for i := 0; i < testCases; i++ {
            fmt.Scan(&coordinate)
            process(coordinate)
        }
    }
}
