'use strict'

const crypto = require('crypto')

exports.genIntN = (min, max) => {
  return crypto.randomInt(min, max + 1)
}

exports.genInt = () => {
  return this.genIntN(1, 2_147_483_647)
}

exports.genEven = () => {
  return this.genInt() >>> 1 << 1
}

exports.genOdd = () => {
  return this.genInt() | 1
}
