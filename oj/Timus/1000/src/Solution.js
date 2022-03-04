"use strict"

const readline = require("readline")

const input = readline.createInterface({
    input: process.stdin,
    output: process.stdout
})

if (module === require.main) {
    input.on("line", (line) => {
        let pair = line.replace(/(^\s*)|(\s*$)/g, '').split(/ +/)
        let a = Number(pair[0])
        let b = Number(pair[1])
        console.log(a + b)
    })
}
