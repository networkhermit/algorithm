import * as GreatestCommonDivisor from './GreatestCommonDivisor.js';

export const reduceToBinaryGCD = (m, n) => {
  return GreatestCommonDivisor.iterativeBinaryGCD(m, n) === 1;
};

export const reduceToEuclidean = (m, n) => {
  return GreatestCommonDivisor.iterativeEuclidean(m, n) === 1;
};
