"use strict"

const BinarySearch = require("vac.fun/algorithms/search/BinarySearch")
const LinearSearch = require("vac.fun/algorithms/search/LinearSearch")

exports.binarySearch = (arr, key) => {
    return BinarySearch.find(arr, key)
}

exports.linearSearch = (arr, key) => {
    return LinearSearch.find(arr, key)
}
