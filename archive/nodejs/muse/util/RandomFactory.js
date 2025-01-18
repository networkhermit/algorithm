import * as crypto from 'node:crypto';

export const genIntN = (min, max) => {
  return crypto.randomInt(min, max + 1);
};

export const genInt = () => {
  return genIntN(1, 2_147_483_647);
};

export const genEven = () => {
  return (genInt() >>> 1) << 1;
};

export const genOdd = () => {
  return genInt() | 1;
};
