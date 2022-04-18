package muse.algorithms.number_theory;

public final class Primality {

  private Primality() {}

  public static boolean isPrime(long n) {
    if (n < 2) {
      return false;
    }
    if ((n & 1) == 0 && n != 2) {
      return false;
    }

    for (int i = 3, bound = (int) Math.sqrt((double) n) + 1; i < bound; i += 2) {
      if (n % i == 0) {
        return false;
      }
    }

    return true;
  }

  public static boolean isComposite(long n) {
    return n > 1 && !isPrime(n);
  }
}
