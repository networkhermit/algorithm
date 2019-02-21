"use strict"

const BinarySearch = require("fun/vac/algorithms/search/BinarySearch")
const LinearSearch = require("fun/vac/algorithms/search/LinearSearch")

exports.binarySearch = (arr, key) => {
    return BinarySearch.find(arr, key)
}

exports.linearSearch = (arr, key) => {
    return LinearSearch.find(arr, key)
}
