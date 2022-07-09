'use strict'

const RandomFactory = require('muse/util/RandomFactory')
const Sequences = require('muse/util/Sequences')

exports.packIncreasing = (arr) => {
  arr[0] = RandomFactory.genIntN(1, 3)
  for (let i = 1, length = arr.length; i < length; i++) {
    arr[i] = arr[i - 1] + RandomFactory.genIntN(1, 3)
  }
}

exports.packRandom = (arr) => {
  for (let i = 0, length = arr.length; i < length; i++) {
    arr[i] = RandomFactory.genInt()
  }
}

exports.packDecreasing = (arr) => {
  this.packIncreasing(arr)
  Sequences.reverse(arr)
}
