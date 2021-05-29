"use strict"

const BubbleSort    = require("vac.fun/algorithms/sorting/BubbleSort")
const InsertionSort = require("vac.fun/algorithms/sorting/InsertionSort")
const MergeSort     = require("vac.fun/algorithms/sorting/MergeSort")
const QuickSort     = require("vac.fun/algorithms/sorting/QuickSort")
const SelectionSort = require("vac.fun/algorithms/sorting/SelectionSort")

exports.bubbleSort = (arr) => {
    BubbleSort.sort(arr)
}

exports.insertionSort = (arr) => {
    InsertionSort.sort(arr)
}

exports.mergeSort = (arr) => {
    MergeSort.sort(arr)
}

exports.quickSort = (arr) => {
    QuickSort.sort(arr)
}

exports.selectionSort = (arr) => {
    SelectionSort.sort(arr)
}
