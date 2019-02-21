"use strict"

const BubbleSort    = require("fun/vac/algorithms/sorting/BubbleSort")
const InsertionSort = require("fun/vac/algorithms/sorting/InsertionSort")
const MergeSort     = require("fun/vac/algorithms/sorting/MergeSort")
const QuickSort     = require("fun/vac/algorithms/sorting/QuickSort")
const SelectionSort = require("fun/vac/algorithms/sorting/SelectionSort")

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
