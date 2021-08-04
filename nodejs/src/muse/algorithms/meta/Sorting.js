"use strict"

const BubbleSort    = require("muse/algorithms/sorting/BubbleSort")
const InsertionSort = require("muse/algorithms/sorting/InsertionSort")
const MergeSort     = require("muse/algorithms/sorting/MergeSort")
const QuickSort     = require("muse/algorithms/sorting/QuickSort")
const SelectionSort = require("muse/algorithms/sorting/SelectionSort")

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
