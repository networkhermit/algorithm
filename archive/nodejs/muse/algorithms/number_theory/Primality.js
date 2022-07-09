'use strict'

exports.isPrime = (n) => {
  if (n < 2) {
    return false
  }
  if ((n & 1) === 0 && n !== 2) {
    return false
  }

  for (let i = 3, bound = Math.floor(Math.sqrt(n)) + 1; i < bound; i += 2) {
    if (n % i === 0) {
      return false
    }
  }

  return true
}

exports.isComposite = (n) => {
  return n > 1 && !this.isPrime(n)
}
