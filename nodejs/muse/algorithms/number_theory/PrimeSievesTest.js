"use strict"

const Primality   = require("muse/algorithms/number_theory/Primality")
const PrimeSieves = require("muse/algorithms/number_theory/PrimeSieves")
const TestRunner  = require("muse/util/TestRunner")

const testPrimeSieves = () => {
    let sample = [
        [      0,     0],
        [      1,     0],
        [    180,    41],
        [     26,     9],
        [    350,    70],
        [    521,    98],
        [    307,    63],
        [    163,    38],
        [    181,    42],
        [    372,    73],
        [    421,    82],
        [    241,    53],
        [    229,    50],
        [    169,    39],
        [    442,    85],
        [    445,    86],
        [    338,    68],
        [    306,    62],
        [     56,    16],
        [    270,    57],
        [    499,    95],
        [    173,    40],
        [     18,     7],
        [     43,    14],
        [    218,    47],
        [    122,    30],
        [    265,    56],
        [    467,    91],
        [    381,    75],
        [     11,     5],
        [    409,    80],
        [    233,    51],
        [   1119,   187],
        [   6976,   896],
        [   4355,   594],
        [   7757,   984],
        [   3467,   486],
        [   5824,   764],
        [   7724,   980],
        [   1470,   232],
        [   7894,   997],
        [   6788,   873],
        [   7873,   994],
        [   6874,   885],
        [   3554,   497],
        [   4137,   569],
        [   7878,   995],
        [   3390,   477],
        [   2141,   323],
        [   4642,   626],
        [   2718,   396],
        [   7126,   912],
        [   1022,   172],
        [   5651,   743],
        [   4822,   649],
        [   2022,   306],
        [   6449,   837],
        [   2153,   325],
        [   1323,   216],
        [   6960,   893],
        [   2837,   412],
        [   4212,   576],
        [   2179,   327],
        [   1033,   174],
        [  33023,  3540],
        [  65798,  6572],
        [  39761,  4179],
        [  12379,  1478],
        [  25339,  2794],
        [  50093,  5143],
        [  98641,  9471],
        [  84959,  8272],
        [  90697,  8778],
        [  20898,  2350],
        [  63079,  6326],
        [  18369,  2104],
        [  23459,  2611],
        [  55337,  5620],
        [  55057,  5596],
        [  94463,  9111],
        [  22769,  2544],
        [  22501,  2516],
        [  68369,  6799],
        [  65315,  6524],
        [  72893,  7206],
        [  44969,  4672],
        [  69511,  6902],
        [  24371,  2705],
        [  32579,  3497],
        [  26017,  2862],
        [  84309,  8217],
        [ 100043,  9595],
        [  86861,  8440],
        [  79355,  7776],
        [  14445,  1693],
        [  95418,  9197],
        [1209025, 93544],
        [ 526733, 43583],
        [ 832339, 66339],
        [ 350593, 30017],
        [ 616581, 50358],
        [ 787123, 62995],
        [ 474923, 39612],
        [ 320448, 27647],
        [ 617049, 50394],
        [ 232333, 20622],
        [ 651727, 52944],
        [ 524873, 43427],
        [1008852, 79161],
        [ 187573, 16970],
        [ 432235, 36339],
        [ 714479, 57629],
        [ 669287, 54254],
        [ 633133, 51574],
        [ 964012, 75920],
        [ 663127, 53800],
        [1257461, 96995],
        [ 172171, 15681],
        [1112392, 86634],
        [ 912533, 72211],
        [ 409529, 34574],
        [ 691575, 55914],
        [ 742768, 59730],
        [ 479430, 39964],
        [ 108289, 10303],
        [1228187, 94920],
        [ 709431, 57257],
        [1294061, 99610],
    ]

    let arr = null

    for (let i = 0, size = sample.length; i < size; i++) {
        arr = PrimeSieves.sieveOfEratosthenes(sample[i][0])

        if (arr.length !== sample[i][1]) {
            return false
        }

        for (let v of arr) {
            if (!Primality.isPrime(v)) {
                return false
            }
        }
    }

    return true
}

if (module === require.main) {
    TestRunner.pick(testPrimeSieves)
}
