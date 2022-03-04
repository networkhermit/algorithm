"use strict"

if (module === require.main) {
    console.log("When’s Weekend?")

    let today = new Date().getDay()

    switch (today) {
    case 0:
        console.log("Today.")
        break
    case 6:
        console.log("Today.")
        break
    case 5:
        console.log("Tomorrow.")
        break
    case 4:
        console.log("In two days.")
        break
    default:
        console.log("Too far away.")
    }
}
