import * as GreatestCommonDivisor from './GreatestCommonDivisor.js';

export const reduceToBinaryGCD = (m, n) => {
  if (m < 0) {
    m = -m;
  }
  if (n < 0) {
    n = -n;
  }

  if (m === 0 || n === 0) {
    return 0;
  }
  return (m / GreatestCommonDivisor.iterativeBinaryGCD(m, n)) * n;
};

export const reduceToEuclidean = (m, n) => {
  if (m < 0) {
    m = -m;
  }
  if (n < 0) {
    n = -n;
  }

  if (m === 0 || n === 0) {
    return 0;
  }
  return (m / GreatestCommonDivisor.iterativeEuclidean(m, n)) * n;
};
