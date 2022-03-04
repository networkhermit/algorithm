"use strict"

const fizzBuzz = (n) => {
    let word = ""

    if (n % 3 === 0) {
        word += "Fizz"
    }
    if (n % 5 === 0) {
        word += "Buzz"
    }

    if (word.length === 0) {
        word = String(n)
    }

    return word
}

if (module === require.main) {
    for (let i = 1; i <= 100; i++) {
        console.log(fizzBuzz(i))
    }
}
