'use strict';

exports.sieveOfEratosthenes = (n) => {
  if (n < 2) {
    return new Array(0);
  }

  const size = (n + 1) >>> 1;

  const arr = new Array(size);
  arr.fill(false);

  let numPrimes = size;
  for (let i = 3, bound = Math.floor(Math.sqrt(n)) + 1; i < bound; i += 2) {
    if (arr[i >>> 1]) {
      continue;
    }
    for (let j = i * i; j <= n; j += i << 1) {
      if (!arr[j >>> 1]) {
        arr[j >>> 1] = true;
        numPrimes--;
      }
    }
  }

  const primes = new Array(numPrimes);

  primes[0] = 2;

  let curr = 1;
  for (let i = 3, bound = n + 1; i < bound; i += 2) {
    if (!arr[i >>> 1]) {
      primes[curr] = i;
      curr++;
    }
  }

  return primes;
};
