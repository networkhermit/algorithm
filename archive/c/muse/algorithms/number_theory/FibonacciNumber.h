#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_FIBONACCI_NUMBER_H
#define MUSE_ALGORITHMS_NUMBER_THEORY_FIBONACCI_NUMBER_H

long FibonacciNumber_iterativeProcedure(long n) {
  long sign = 1;

  if (n < 0) {
    if ((n & 1) == 0) {
      sign = -1;
    }
    n = -n;
  }

  if (n < 2) {
    return n;
  }

  long prev = 0;
  long curr = 1;

  long next;
  for (long i = 2; i <= n; i++) {
    next = prev + curr;
    prev = curr;
    curr = next;
  }

  return sign * curr;
}

long FibonacciNumber_recursiveProcedure(long n) {
  if (n < 0) {
    if ((n & 1) == 0) {
      return -FibonacciNumber_recursiveProcedure(-n);
    }
    return FibonacciNumber_recursiveProcedure(-n);
  }
  if (n < 2) {
    return n;
  }
  return FibonacciNumber_recursiveProcedure(n - 2) +
         FibonacciNumber_recursiveProcedure(n - 1);
}

#endif
