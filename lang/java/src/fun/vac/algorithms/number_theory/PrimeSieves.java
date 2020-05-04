package fun.vac.algorithms.number_theory;

public class PrimeSieves {

    private PrimeSieves() {}

    public static int[] sieveOfEratosthenes(int n) {
        if (n < 2) {
            return new int[0];
        }

        int size = (n + 1) >>> 1;

        boolean[] arr = new boolean[size];

        int numPrimes = size;
        for (int i = 3, bound = (int) Math.sqrt((double) n) + 1; i < bound; i += 2) {
            if (arr[i >>> 1]) {
                continue;
            }
            for (int j = i * i; j <= n; j += i << 1) {
                if (!arr[j >>> 1]) {
                    arr[j >>> 1] = true;
                    numPrimes--;
                }
            }
        }

        int[] primes = new int[numPrimes];

        primes[0] = 2;

        int curr = 1;
        for (int i = 3, bound = n + 1; i < bound; i += 2) {
            if (!arr[i >>> 1]) {
                primes[curr] = i;
                curr++;
            }
        }

        return primes;
    }
}
